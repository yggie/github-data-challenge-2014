package models

type CommitAuthor struct {
	Name  string
	Email string
}

type Commit struct {
	Sha      string
	Distinct bool
	Message  string
	Author   *CommitAuthor
}
