package models

type Event interface {
	Id() int64
	Type() string
	CreatedAt() string
	Actor() *Actor
	Repo() *Repo
}
