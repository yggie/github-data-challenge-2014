package models

type CommitAuthor struct {
	Name  string
	Email string
}

type Commit struct {
	sha      string
	distinct bool
	message  string
	Author   CommitAuthor
}

func NewCommit(raw interface{}) Commit {
	data := raw.(map[string]interface{})
	author := data["author"].(map[string]interface{})
	return Commit{
		sha:      data["sha"].(string),
		message:  data["message"].(string),
		distinct: data["distinct"].(bool),
		Author: CommitAuthor{
			Name:  author["name"].(string),
			Email: author["email"].(string),
		},
	}
}

func NewCommits(data interface{}) []Commit {
	array := data.([]interface{})
	commits := make([]Commit, len(array))
	for index, element := range array {
		commits[index] = NewCommit(element)
	}

	return commits
}
