package models

// A simple representation of a GitHub User present in the GitHub API Event
// objects
type User struct {
	Id         int64
	Login      string
	GravatarId string
	AvatarUrl  string
}
