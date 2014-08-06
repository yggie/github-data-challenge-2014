package models

// Constructs an appropriate Event struct based on the type provided
func NewEvent(raw interface{}) (Event, bool) {
	object := raw.(map[string]interface{})
	switch object["type"] {
	case "PushEvent":
		return NewPushEvent(object), true
	}

	return nil, false
}

// Shared interface for a GitHub Event
type Event interface {
	Id() int64
	Type() string
	CreatedAt() string
	User() *User
	Repository() *Repository
}
