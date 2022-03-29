package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Init() (*sql.DB, error) {
	dsn := os.Getenv(("MYSQL_USER")) + ":" + os.Getenv(("MYSQL_PASSWORD")) + "@tcp(" + os.Getenv(("MYSQL_HOST")) + ":" + os.Getenv(("MYSQL_PORT")) + ")/" + os.Getenv(("MYSQL_DATABASE")) + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr := sql.Open("mysql", dsn)
	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}
