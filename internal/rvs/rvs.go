package rvs

import (
	"fmt"
	"os"

	"github.com/osm/rvspub/internal/buffer"
	"github.com/osm/rvspub/internal/event"
)

func FromFile(path string) ([]event.Event, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %w", path, err)
	}

	events, err := parse(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rvs data: %w", err)
	}

	return events, nil
}

func parse(buf []byte) ([]event.Event, error) {
	b := buffer.New(buf)

	// TODO: Figure out what these bytes are.
	if err := b.SkipBytes(128); err != nil {
		return nil, fmt.Errorf("failed to read rvs data: %v", err)
	}

	var events []event.Event

	for {
		if b.IsEnd() {
			break
		}

		event, err := event.FromBuffer(b)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
