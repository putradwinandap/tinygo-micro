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
			"omg":    true,
			"number": 123,
		}).Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	return db, nil
}
