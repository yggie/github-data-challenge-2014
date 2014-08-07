package models

// A shared representation of a GitHub API Event object, exposes a simple
// interface to get the properties of the Event
type Event struct {
	Id int64
	// The Event type, can be one of the following:
	// CommitCommentEvent, CreateEvent, DeleteEvent, DeploymentEvent,
	// DeploymentStatusEvent, DownloadEvent, FollowEvent, ForkEvent, ForkApplyEvent,
	// GistEvent, GollumEvent, IssueCommentEvent, IssueEvent, MemberEvent,
	// PageBuildEvent, PublicEvent, PullRequestEvent, PullRequestReviewCommentEvent,
	// PushEvent, ReleaseEvent, StatusEvent, TeamAddEvent, WatchEvent
	EventType  string
	CreatedAt  string
	User       *User
	Repository *Repository
}
