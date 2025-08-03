package redis

import (
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func NewClient(dsn string) (*redis.Client, error) {
	//url := "redis://user:password@localhost:6379/0?protocol=3"
	url := dsn
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to parse Redis URL")
		return nil, err
	}
	return redis.NewClient(opts), nil
}
