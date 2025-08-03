package shorturl

import (
	"url/internal/domain/entity"
	"url/internal/domain/iface"
)

type FindByIDShortURLUseCase struct {
	repo iface.ShortURLRepository
}

func NewFindByIDShortURLUseCase(repo iface.ShortURLRepository) *FindByIDShortURLUseCase {
	return &FindByIDShortURLUseCase{
		repo: repo,
	}
}

func (uc *FindByIDShortURLUseCase) Execute(id int) (entity.ShortUrl, error) {
	shortURL, err := uc.repo.FindByID(id)
	if err != nil {
		return entity.ShortUrl{}, err
	}

	return shortURL, nil
}
