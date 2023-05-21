package db

import (
	"database/sql"
	"fmt"
	"os"
)

// TODO:他のディレクトリに移動
func NewPostgresDB() (*sql.DB, error) {
	host := os.Getenv("PSQL_HOST")
	dbname := os.Getenv("PSQL_DBNAME")
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASS")
	sslmode := os.Getenv("PSQL_SSLMODE")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", host, dbname, user, password, sslmode))
	return db, err
}

func NewMySQLDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s&charset=utf8&parseTime=true", dsn))
	return db, err
}
