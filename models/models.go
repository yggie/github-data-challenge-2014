package models

import "encoding/json"

func ParseEvents(data []byte) []Event {
	var rawEvents []interface{}
	json.Unmarshal(data, &rawEvents)

	events := make([]Event, len(rawEvents))
	count := 0
	for _, element := range rawEvents {
		event, ok := NewEvent(element)
		if ok {
			events[count] = event
			count++
		}
	}
	return events[0:count]
}
