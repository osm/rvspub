package csv

import (
	"encoding/csv"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/osm/rvspub/internal/event"
)

func Format(events []event.Event, fields []string) (string, error) {
	var b strings.Builder
	writer := csv.NewWriter(&b)

	if err := writer.Write(fields); err != nil {
		return "", fmt.Errorf("failed to write csv header: %w", err)
	}

	t := reflect.TypeOf(event.Event{})
	indexes := make([]int, len(fields))
	for i, field := range fields {
		for j := 0; j < t.NumField(); j++ {
			if t.Field(j).Tag.Get("json") == field {
				indexes[i] = j
				break
			}
		}
	}

	for _, e := range events {
		v := reflect.ValueOf(e)
		row := make([]string, len(fields))

		for i, idx := range indexes {
			row[i] = formatValue(v.Field(idx))
		}

		if err := writer.Write(row); err != nil {
			return "", fmt.Errorf("failed to write csv data: %w", err)
		}
	}

	return b.String(), nil
}

func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Struct:
		if t, ok := v.Interface().(time.Time); ok {
			return t.Format(time.RFC3339)
		}
	}

	if ip, ok := v.Interface().(net.IP); ok {
		return ip.String()
	}

	return fmt.Sprintf("%v", v.Interface())
}
