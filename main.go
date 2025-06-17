package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osm/rvspub/internal/fields"
	"github.com/osm/rvspub/internal/format/csv"
	"github.com/osm/rvspub/internal/format/json"
	"github.com/osm/rvspub/internal/format/text"
	"github.com/osm/rvspub/internal/rvs"
)

func main() {
	dataFile := flag.String("data-file", "", "Path to the data file")
	format := flag.String("format", "csv", "Output format (csv, json, text)")
	fieldsRaw := flag.String("fields", "all", "Comma separated list of fields")
	flag.Parse()

	if *dataFile == "" {
		fmt.Fprintf(os.Stderr, "-data-file is required\n")
		os.Exit(1)
	}
	if *format != "csv" && *format != "json" && *format != "text" {
		fmt.Fprintf(os.Stderr, "-format %q is unknown\n", *format)
		os.Exit(1)
	}

	fields, err := fields.Parse(*fieldsRaw)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse -fields: %v\n", err)
		os.Exit(1)
	}

	events, err := rvs.FromFile(*dataFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse RVS data: %v\n", err)
		os.Exit(1)
	}

	var output string
	switch *format {
	case "csv":
		output, err = csv.Format(events, fields)
	case "json":
		output, err = json.Format(events, fields)
	default:
		output, err = text.Format(events, fields)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format %v data: %v\n", *format, err)
		os.Exit(1)
	}
	fmt.Println(output)
}
