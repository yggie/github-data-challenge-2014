package models

// A simple representation of a GitHub User present in the GitHub API Event
// objects
type User struct {
	id         int64
	login      string
	gravatarId string
	avatarUrl  string
}

func NewUser(raw interface{}) User {
	data := raw.(map[string]interface{})
	return User{
		id:         int64(data["id"].(float64) + 0.5),
		login:      data["login"].(string),
		gravatarId: data["gravatar_id"].(string),
		avatarUrl:  data["avatar_url"].(string),
	}
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Login() string {
	return u.login
}

func (u *User) GravatarId() string {
	return u.gravatarId
}

func (u *User) AvatarUrl() string {
	return u.avatarUrl
}
