package mining

import (
	"encoding/json"
	"strconv"

	"github.com/yggie/github-data-challenge-2014/models"
)

type EventsResult struct {
	PushEvents []*models.PushEvent
}

func (r *EventsResult) AddPushEvent(event *models.PushEvent) {
	r.PushEvents = append(r.PushEvents, event)
}

func ParseEvents(data []byte) *EventsResult {
	var rawEvents []interface{}
	json.Unmarshal(data, &rawEvents)

	result := EventsResult{}
	for _, element := range rawEvents {
		elem := element.(map[string]interface{})

		switch elem["type"] {
		case "PushEvent":
			result.AddPushEvent(ToPushEvent(elem))
		}
	}

	return &result
}

func ToEvent(data map[string]interface{}) *models.Event {
	id, err := strconv.ParseInt(data["id"].(string), 10, 64)
	if err != nil {
		panic(err)
	}

	return &models.Event{
		Id:         id,
		EventType:  data["type"].(string),
		CreatedAt:  data["created_at"].(string),
		User:       ToUser(data["actor"].(map[string]interface{})),
		Repository: ToRepository(data["repo"].(map[string]interface{})),
	}
}

func ToPushEvent(data map[string]interface{}) *models.PushEvent {
	payload := data["payload"].(map[string]interface{})

	return &models.PushEvent{
		Event:   ToEvent(data),
		Size:    int(payload["size"].(float64) + 0.5),
		PushId:  int64(payload["push_id"].(float64) + 0.5),
		Commits: ToCommits(payload["commits"].([]interface{})),
	}
}

func ToUser(data map[string]interface{}) *models.User {
	return &models.User{
		Id:         int64(data["id"].(float64) + 0.5),
		Login:      data["login"].(string),
		GravatarId: data["gravatar_id"].(string),
		AvatarUrl:  data["avatar_url"].(string),
	}
}

func ToRepository(data map[string]interface{}) *models.Repository {
	return &models.Repository{
		Id:   int64(data["id"].(float64) + 0.5),
		Name: data["name"].(string),
		Url:  data["url"].(string),
	}
}

func ToCommit(data map[string]interface{}) *models.Commit {
	author := data["author"].(map[string]interface{})
	return &models.Commit{
		Sha:      data["sha"].(string),
		Message:  data["message"].(string),
		Distinct: data["distinct"].(bool),
		Author: &models.CommitAuthor{
			Name:  author["name"].(string),
			Email: author["email"].(string),
		},
	}
}

func ToCommits(data []interface{}) []*models.Commit {
	commits := make([]*models.Commit, len(data))
	for index, element := range data {
		commits[index] = ToCommit(element.(map[string]interface{}))
	}

	return commits
}
