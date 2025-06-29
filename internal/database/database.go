package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driverName string = "postgres"

var dbInstance *sqlx.DB

func Init(host string, user string, password string, port string, dbName string, sslMode string) (*sqlx.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		host,
		user,
		password,
		port,
		dbName,
		sslMode,
	)

	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbInstance = db

	return dbInstance, nil
}
