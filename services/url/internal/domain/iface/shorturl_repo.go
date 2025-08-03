package iface

import (
	"url/internal/domain/entity"
)

type ShortURLRepository interface {
	Save(shortURL entity.ShortUrl) (entity.ShortUrl, error)
	FindByID(id int) (entity.ShortUrl, error)
	FindByShortcode(shortcode string) (entity.ShortUrl, error)
}
