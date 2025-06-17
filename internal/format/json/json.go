package json

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/osm/rvspub/internal/event"
)

func Format(events []event.Event, fields []string) (string, error) {
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

	output := make([]map[string]any, len(events))
	for i, e := range events {
		v := reflect.ValueOf(e)
		m := make(map[string]any, len(fields))
		for j, idx := range indexes {
			m[fields[j]] = v.Field(idx).Interface()
		}
		output[i] = m
	}

	data, err := json.Marshal(output)
	if err != nil {
		return "", fmt.Errorf("failed to marshal events: %v", err)
	}

	return string(data), nil
}
