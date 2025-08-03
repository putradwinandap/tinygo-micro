package dto

type SaveShortURLRequest struct {
	Url string `json:"url" binding:"required"`
}

type FindByIDShortURLRequest struct {
	ID int `json:"id" binding:"required"`
}

type ResolveShortURLRequest struct {
	ShortCode string `uri:"shortcode" binding:"required"`
}
