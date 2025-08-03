package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to connect to the database")
		return nil, err
	}
	return db, nil
}
