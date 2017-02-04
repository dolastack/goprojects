package cms

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=goprojects dbname=goprojects sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &PgStore{
		DB: db,
	}
}
