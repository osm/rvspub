package fields

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/osm/rvspub/internal/event"
)

func Parse(input string) ([]string, error) {
	eventType := reflect.TypeOf(event.Event{})
	jsonFieldSet := make(map[string]struct{})
	for i := 0; i < eventType.NumField(); i++ {
		field := eventType.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" && tag != "-" {
			jsonFieldSet[tag] = struct{}{}
		}
	}

	fields := []string{}
	if input == "all" {
		for f := range jsonFieldSet {
			fields = append(fields, f)
		}
	} else {
		for _, f := range toSlice(input) {
			if _, ok := jsonFieldSet[f]; !ok {
				return nil, fmt.Errorf("failed to find field: %q", f)
			}
			fields = append(fields, f)
		}
	}

	return fields, nil
}

func toSlice(input string) []string {
	var result []string

	for _, part := range strings.Split(input, ",") {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
