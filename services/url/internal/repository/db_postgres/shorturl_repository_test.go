package db_postgres_test

import (
	"testing"
	"time"

	"url/internal/domain/entity"
	"url/internal/repository/db_postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := db_postgres.NewShortURLRepository(db)

	UserID := 1

	expected := entity.ShortUrl{
		ShortCode: "abc123",
		LongUrl:   "http://example.com",
		UserID:    &UserID,
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "short_code", "long_url", "user_id", "created_at"}).
		AddRow(1, expected.ShortCode, expected.LongUrl, expected.UserID, expected.CreatedAt)

	mock.ExpectQuery(`INSERT INTO short_urls`).WithArgs(expected.ShortCode, expected.LongUrl, expected.UserID, expected.CreatedAt).
		WillReturnRows(rows)

	result, err := repo.Save(expected)
	assert.NoError(t, err)
	assert.Equal(t, expected.ShortCode, result.ShortCode)
}

func TestFindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := db_postgres.NewShortURLRepository(db)

	expected := entity.ShortUrl{
		ID:        1,
		ShortCode: "abc123",
		LongUrl:   "http://example.com",
		UserID:    nil,
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "short_code", "long_url", "user_id", "created_at"}).
		AddRow(expected.ID, expected.ShortCode, expected.LongUrl, expected.UserID, expected.CreatedAt)

	mock.ExpectQuery(`SELECT id, short_code, long_url, user_id, created_at FROM short_urls WHERE id = \$1`).
		WithArgs(expected.ID).
		WillReturnRows(rows)

	result, err := repo.FindByID(expected.ID)
	assert.NoError(t, err)
	assert.Equal(t, expected.ShortCode, result.ShortCode)
}
