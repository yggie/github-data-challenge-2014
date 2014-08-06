package models

import "errors"

type UnsupportedEventType error

func NewUnsupportedEventType(eventType string) UnsupportedEventType {
	return errors.New("Unsupported Event Type: [" + eventType + "]")
}
