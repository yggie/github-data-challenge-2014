package models

import "github.com/jmcvetta/neoism"

type Tracked interface {
	Id() int64
}

func CountEvents() int {
	var result interface{}
	query := neoism.CypherQuery{
		Statement: `MATCH (e:Event) RETURN count(e) as count`,
		Result:    &result,
	}

	err := db.Cypher(&query)
	if err != nil {
		panic(err)
	}

	res := result.([]interface{})[0].(map[string]interface{})
	return int(res["count"].(float64) + 0.5)
}

func CheckExists(modelName string, object Tracked) bool {
	var result []interface{}
	query := neoism.CypherQuery{
		Statement: `MATCH (e:` + modelName + ` { id: {id} }) RETURN count(*) as count`,
		Parameters: neoism.Props{
			"id": object.Id(),
		},
		Result: &result,
	}

	err := db.Cypher(&query)
	if err != nil {
		panic(err)
	}

	return int(result[0].(map[string]interface{})["count"].(float64)+0.5) != 0
}

func ClearEvents() error {
	query := neoism.CypherQuery{
		Statement: `MATCH (e:Event) DELETE e`,
	}

	return db.Cypher(&query)
}
