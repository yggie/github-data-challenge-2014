package neo

import (
	"github.com/jmcvetta/neoism"
	"github.com/yggie/github-data-challenge-2014/models"
)

func PersistPushEvent(event *models.PushEvent) error {
	queries := make([]*neoism.CypherQuery, 0)

	if !CheckExists(EVENTS, event.Id) {
		query := neoism.CypherQuery{
			Statement: `CREATE (:` + string(EVENTS) + ` { id: {id}, type: {type}, created_at: {created_at} })`,
			Parameters: neoism.Props{
				"id":         event.Id,
				"type":       event.EventType,
				"created_at": event.CreatedAt,
			},
		}

		queries = append(queries, &query)
	}

	repository := event.Repository
	if !CheckExists(REPOSITORIES, repository.Id) {
		query := &neoism.CypherQuery{
			Statement: `CREATE (:` + string(REPOSITORIES) + ` { id: {id}, name: {name}, url: {url} })`,
			Parameters: neoism.Props{
				"id":   repository.Id,
				"name": repository.Name,
				"url":  repository.Url,
			},
		}

		queries = append(queries, query)

		distribution := repository.LanguageDistribution()
		for key, value := range distribution {
			query = &neoism.CypherQuery{
				Statement: `
					MATCH (r:` + string(REPOSITORIES) + ` { id: {repository_id} })
					CREATE UNIQUE (r)-[:` + string(WRITTEN_IN) + ` { weight: {weight} }]->(:` + string(LANGUAGES) + ` { name: {name} })`,
				Parameters: neoism.Props{
					"repository_id": repository.Id,
					"name":          key,
					"weight":        value,
				},
			}

			queries = append(queries, query)
		}
	}

	user := event.User
	if !CheckExists(USERS, user.Id) {
		query := neoism.CypherQuery{
			Statement: `CREATE (:` + string(USERS) + ` {
				id: {id},
				login: {login},
				gravatar_id: {gravatar_id},
				avatar_url: {avatar_url}
			})`,
			Parameters: neoism.Props{
				"id":          user.Id,
				"login":       user.Login,
				"gravatar_id": user.GravatarId,
				"avatar_url":  user.AvatarUrl,
			},
		}

		queries = append(queries, &query)
	}

	eventToRepository := neoism.CypherQuery{
		Statement: `
			MATCH (r:Repository { id: {repository_id} }), (e:Event { id: {event_id} })
			CREATE UNIQUE (r)-[:HAS_A]->(e)`,
		Parameters: neoism.Props{
			"event_id":      event.Id,
			"repository_id": repository.Id,
		},
	}

	queries = append(queries, &eventToRepository)

	tx, err := db.Begin(queries)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()

		if err2 != nil {
			return err2
		}

		return err
	}

	return nil
}
