package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	strConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
