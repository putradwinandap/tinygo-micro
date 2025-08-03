package db_postgres

import (
	"database/sql"
	"url/internal/domain/entity"
)

type ShortURLRepository struct {
	DB *sql.DB
}

// NewShortURLRepository creates a new instance of ShortURLRepository.
func NewShortURLRepository(db *sql.DB) *ShortURLRepository {
	return &ShortURLRepository{
		DB: db,
	}
}

func (r *ShortURLRepository) Save(shortURL entity.ShortUrl) (entity.ShortUrl, error) {
	query := `INSERT INTO short_urls (short_code, long_url, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, short_code, long_url, user_id, created_at`
	err := r.DB.QueryRow(query, shortURL.ShortCode, shortURL.LongUrl, shortURL.UserID, shortURL.CreatedAt).Scan(&shortURL.ID, &shortURL.ShortCode, &shortURL.LongUrl, &shortURL.UserID, &shortURL.CreatedAt)
	if err != nil {
		return entity.ShortUrl{}, err
	}
	return shortURL, nil
}

func (r *ShortURLRepository) FindByID(id int) (entity.ShortUrl, error) {
	query := `SELECT id, short_code, long_url, user_id, created_at FROM short_urls WHERE id = $1`
	var shortURL entity.ShortUrl
	err := r.DB.QueryRow(query, id).Scan(&shortURL.ID, &shortURL.ShortCode, &shortURL.LongUrl, &shortURL.UserID, &shortURL.CreatedAt)
	if err != nil {
		return entity.ShortUrl{}, err
	}
	return shortURL, nil
}

func (r *ShortURLRepository) FindByShortcode(shortcode string) (entity.ShortUrl, error) {
	query := `SELECT id, short_code, long_url, user_id, created_at FROM short_urls WHERE short_code = $1`
	var shortURL entity.ShortUrl
	err := r.DB.QueryRow(query, shortcode).Scan(&shortURL.ID, &shortURL.ShortCode, &shortURL.LongUrl, &shortURL.UserID, &shortURL.CreatedAt)
	if err != nil {
		return entity.ShortUrl{}, err
	}
	return shortURL, nil
}
