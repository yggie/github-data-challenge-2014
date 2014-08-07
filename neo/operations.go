package neo

import "github.com/jmcvetta/neoism"

func Count(modelType Type) int {
	var statement string
	var result interface{}

	if modelType == ALL {
		statement = `MATCH (n) RETURN count(n) as count`
	} else {
		statement = `MATCH (:` + class[modelType] + `) RETURN count(*) as count`
	}

	query := neoism.CypherQuery{
		Statement: statement,
		Result:    &result,
	}

	err := db.Cypher(&query)
	if err != nil {
		panic(err)
	}

	res := result.([]interface{})[0].(map[string]interface{})
	return int(res["count"].(float64) + 0.5)
}

func CheckExists(modelType Type, id int64) bool {
	var result []interface{}
	query := neoism.CypherQuery{
		Statement: `MATCH (e:` + class[modelType] + ` { id: {id} }) RETURN count(*) as count`,
		Parameters: neoism.Props{
			"id": id,
		},
		Result: &result,
	}

	err := db.Cypher(&query)
	if err != nil {
		panic(err)
	}

	return int(result[0].(map[string]interface{})["count"].(float64)+0.5) != 0
}

func Clear(modelType Type) error {
	var statement string
	if modelType == ALL {
		statement = `MATCH (n)-[r]-() DELETE n, r`
	} else {
		statement = `MATCH (e:` + class[modelType] + `) DELETE e`
	}

	return db.Cypher(&neoism.CypherQuery{
		Statement: statement,
	})
}
