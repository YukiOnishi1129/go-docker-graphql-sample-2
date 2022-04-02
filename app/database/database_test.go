package database

import (
	"database/sql"
	"testing"
)

func TestInit(t *testing.T) {
	//	connect to db
	db, dbErr := Init()
	if dbErr != nil {
		panic(dbErr)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
}
