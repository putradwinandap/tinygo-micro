package shorturl

import (
	"url/internal/domain/entity"
	"url/internal/domain/iface"
)

type SaveShortURLUseCase struct {
	repo iface.ShortURLRepository
}

func NewSaveShortURLUseCase(repo iface.ShortURLRepository) *SaveShortURLUseCase {
	return &SaveShortURLUseCase{
		repo: repo,
	}
}

func (uc *SaveShortURLUseCase) Execute(shortURL entity.ShortUrl) (entity.ShortUrl, error) {
	save, err := uc.repo.Save(shortURL)
	if err != nil {
		return entity.ShortUrl{}, err
	}
	return save, nil
}
