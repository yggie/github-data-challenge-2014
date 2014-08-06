package models

import (
	"fmt"
	"strings"

	"github.com/jmcvetta/neoism"
)

func Persist(event Event) error {
	var changed bool
	statements := make([]string, 0)
	switch event.Type() {
	case "PushEvent":
		statements, changed = persistPushEvent(statements, event.(*PushEvent))
	default:
		return NewUnsupportedEventType(event.Type())
	}

	if !changed {
		return nil
	}

	query := neoism.CypherQuery{
		Statement: "CREATE " + strings.Join(statements, ","),
	}

	return db.Cypher(&query)
}

func persistPushEvent(statements []string, event *PushEvent) ([]string, bool) {
	if CheckExists("Event", event) {
		return statements, false
	}

	s := fmt.Sprintf(`(e:Event { id: %d, type: "%s", created_at: "%s"})`, event.id, event.eventType, event.createdAt)
	statements = append(statements, s)

	statements, changed := persistRepository(statements, &event.repository)

	return statements, changed
}

func persistRepository(statements []string, repository *Repository) ([]string, bool) {
	if CheckExists("Repository", repository) {
		return statements, false
	}

	return statements, true
}
