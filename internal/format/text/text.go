package text

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/osm/rvspub/internal/event"
)

func Format(events []event.Event, fields []string) (string, error) {
	var b strings.Builder

	for idx, e := range events {
		v := reflect.ValueOf(e)
		t := reflect.TypeOf(e)

		for i, field := range fields {
			for j := 0; j < t.NumField(); j++ {
				if t.Field(j).Tag.Get("json") == field {
					fmt.Fprint(&b, v.Field(j).Interface())
					break
				}
			}

			if i < len(fields)-1 {
				b.WriteString(" ")
			}
		}

		if idx < len(events)-1 {
			b.WriteString("\n")
		}
	}

	return b.String(), nil
}
