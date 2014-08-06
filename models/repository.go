package models

// Constructs a new GitHub repository struct
func NewRepository(raw interface{}) Repository {
	data := raw.(map[string]interface{})
	return Repository{
		id:   int64(data["id"].(float64) + 0.5),
		name: data["name"].(string),
		url:  data["url"].(string),
	}
}

// A GitHub repository
type Repository struct {
	id   int64
	name string
	url  string
}

func (r *Repository) Id() int64 {
	return r.id
}

func (r *Repository) Name() string {
	return r.name
}

func (r *Repository) Url() string {
	return r.url
}
