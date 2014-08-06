package models

// Constructs a new PushEvent given a parsed JSON map object
func NewPushEvent(raw interface{}) *PushEvent {
	object := raw.(map[string]interface{})
	payload := object["payload"].(map[string]interface{})

	return &PushEvent{
		BasicEvent: NewBasicEvent(object),
		size:       int(payload["size"].(float64) + 0.5),
		pushId:     int64(payload["push_id"].(float64) + 0.5),
		commits:    NewCommits(payload["commits"]),
	}
}

// A GitHub Push Event
type PushEvent struct {
	BasicEvent
	size         int
	distinctSize int
	pushId       int64
	commits      []Commit
}

func (e *PushEvent) HeadCommit() Commit {
	return e.commits[e.size-1]
}

func (e *PushEvent) CommitCount() int {
	return e.size
}
