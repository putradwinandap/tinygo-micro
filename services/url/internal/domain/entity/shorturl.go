package entity

import "time"

type ShortUrl struct {
	ID        int       `json:"id"`
	ShortCode string    `json:"short_code"`
	LongUrl   string    `json:"long_url"`
	UserID    *int      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
