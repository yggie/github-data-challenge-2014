package models

type Repo struct {
	Id   int64
	Name string
	Url  string
}

func NewRepo(raw interface{}) Repo {
	data := raw.(map[string]interface{})
	return Repo{
		Id:   int64(data["id"].(float64) + 0.5),
		Name: data["name"].(string),
		Url:  data["url"].(string),
	}
}
