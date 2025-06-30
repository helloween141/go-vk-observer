package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driverName string = "postgres"

var dbInstance *sqlx.DB

func Init(dsn string) (*sqlx.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

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
