package shorturl

import (
	"context"

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

	err = uc.cache.Incr(ctx, shortURL.ShortCode)
	if err != nil {
		log.Warn("Failed to increment view count for shortcode:", shortURL.ShortCode, "Error:", err)
	}

	err = uc.rabbit.Publish("shorturl.resolved", shortURL)
	if err != nil {
		log.Warn("Failed to publish event for resolved short URL:", err)
	}

	return shortURL, nil
}
