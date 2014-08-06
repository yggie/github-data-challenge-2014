package models

import (
	"encoding/json"

	"github.com/jmcvetta/neoism"
)

var db *neoism.Database

func init() {
	var err error
	db, err = neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		panic(err)
	}
}

func DB() *neoism.Database {
	return db
}

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
