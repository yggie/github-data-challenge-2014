package models

import "strconv"

// A shared representation of a GitHub API Event object, exposes a simple
// interface to get the properties of the Event
type BasicEvent struct {
	id        int64
	eventType string
	createdAt string
	actor     Actor
	repo      Repo
}

// Constructs a BasicEvent from a JSON object
func NewBasicEvent(raw interface{}) BasicEvent {
	data := raw.(map[string]interface{})
	id, err := strconv.ParseInt(data["id"].(string), 10, 64)
	if err != nil {
		panic(err)
	}

	return BasicEvent{
		id:        id,
		eventType: data["type"].(string),
		createdAt: data["created_at"].(string),
		actor:     NewActor(data["actor"]),
		repo:      NewRepo(data["repo"]),
	}
}

// The Event ID
func (e *BasicEvent) Id() int64 {
	return e.id
}

// The Event type, can be one of the following:
// CommitCommentEvent, CreateEvent, DeleteEvent, DeploymentEvent,
// DeploymentStatusEvent, DownloadEvent, FollowEvent, ForkEvent, ForkApplyEvent,
// GistEvent, GollumEvent, IssueCommentEvent, IssueEvent, MemberEvent,
// PageBuildEvent, PublicEvent, PullRequestEvent, PullRequestReviewCommentEvent,
// PushEvent, ReleaseEvent, StatusEvent, TeamAddEvent, WatchEvent
func (e *BasicEvent) Type() string {
	return e.eventType
}

// The date the event was created
func (e *BasicEvent) CreatedAt() string {
	return e.createdAt
}

// The initiator of the event (a GitHub user)
func (e *BasicEvent) Actor() *Actor {
	return &e.actor
}

// The repository affected by the event
func (e *BasicEvent) Repo() *Repo {
	return &e.repo
}
