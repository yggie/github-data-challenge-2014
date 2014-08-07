package neo

import "github.com/jmcvetta/neoism"

var db *neoism.Database

func init() {
	var err error
	db, err = neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		panic(err)
	}
}

func DB() *neoism.Database {
	return db
}
