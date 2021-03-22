package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB_USER")+":"+os.Getenv("MYSQL_DB_PASS")+"@("+os.Getenv("MYSQL_DB_HOST")+":"+os.Getenv("MYSQL_DB_PORT")+")/"+os.Getenv("MYSQL_DB_NAME")+"")
	if err != nil {
		return nil, err
	}
	return db, nil
}
