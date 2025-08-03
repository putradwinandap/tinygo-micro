package shorturl

import (
	"context"
	"fmt"

	global_iface "github.com/putradwinandap/tinygo-micro/shared-lib/message_broker/iface"

	"url/internal/domain/entity"
	"url/internal/domain/iface"

	cacheiface "github.com/putradwinandap/tinygo-micro/shared-lib/cache/iface"
	log "github.com/sirupsen/logrus"
)

type ResolveShortURLUseCase struct {
	repo   iface.ShortURLRepository
	rabbit global_iface.EventPublisher
	cache  cacheiface.Counter
}

func NewResolveShortURLUseCase(repo iface.ShortURLRepository, rabbitPub global_iface.EventPublisher, cache cacheiface.Counter) *ResolveShortURLUseCase {
	return &ResolveShortURLUseCase{
		repo:   repo,
		rabbit: rabbitPub,
		cache:  cache,
	}
}

func (uc *ResolveShortURLUseCase) Execute(shortcode string) (entity.ShortUrl, error) {
	shortURL, err := uc.repo.FindByShortcode(shortcode)
	if err != nil {
		return entity.ShortUrl{}, err
	}

	ctx := context.Background()
	key := fmt.Sprintf("shorturl_visit:%d", shortURL.ID)
	err = uc.cache.Incr(ctx, key)
	if err != nil {
		log.WithFields(log.Fields{
			"key":   key,
			"error": err,
		}).Warn("Error incrementing view count in cache")
	}

	err = uc.rabbit.Publish("shorturl.resolved", shortURL)
	if err != nil {

		log.WithFields(log.Fields{
			"shortcode": shortcode,
			"error":     err,
		}).Warn("Failed to publish resolved short URL event")

	}

	return shortURL, nil
}
