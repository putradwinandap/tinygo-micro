package shorturl

import (
	"url/internal/domain/entity"
	"url/internal/dto"
	usecase "url/internal/usecase/shorturl"
	"url/internal/utils"

	"github.com/gin-gonic/gin"
)

type SaveShortURLHandler struct {
	useCase *usecase.SaveShortURLUseCase
}

func NewSaveShortURLHandler(useCase *usecase.SaveShortURLUseCase) *SaveShortURLHandler {
	return &SaveShortURLHandler{
		useCase: useCase,
	}
}

func (h *SaveShortURLHandler) Handle(c *gin.Context) {
	var request dto.SaveShortURLRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	shortURL := entity.ShortUrl{
		LongUrl:   request.Url,
		ShortCode: utils.GenerateShortCode(8),
	}

	shortURL, err = h.useCase.Execute(shortURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save URL"})
		return
	}

	c.JSON(200, gin.H{"short_url": shortURL.ShortCode})
}
