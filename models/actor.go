package models

// A simple representation of a GitHub User present in the GitHub API Event
// objects
type Actor struct {
	Id         int64
	Login      string
	GravatarId string
	AvatarUrl  string
}

func NewActor(raw interface{}) Actor {
	data := raw.(map[string]interface{})
	return Actor{
		Id:         int64(data["id"].(float64) + 0.5),
		Login:      data["login"].(string),
		GravatarId: data["gravatar_id"].(string),
		AvatarUrl:  data["avatar_url"].(string),
	}
}
