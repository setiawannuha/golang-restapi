package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Postgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://"+os.Getenv("PG_DB_USER")+":"+os.Getenv("PG_DB_PASS")+"@"+os.Getenv("PG_DB_HOST")+":"+os.Getenv("PG_DB_PORT")+"/"+os.Getenv("PG_DB_NAME")+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
