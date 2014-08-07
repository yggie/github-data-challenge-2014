package models

// A GitHub Push Event
type PushEvent struct {
	*Event
	Size         int
	DistinctSize int
	PushId       int64
	Commits      []*Commit
}

func (e *PushEvent) HeadCommit() *Commit {
	return e.Commits[e.Size-1]
}
