package shorturl

import (
	"url/internal/dto"
	usecase "url/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
)

type FindByIDShortURLHandler struct {
	useCase *usecase.FindByIDShortURLUseCase
}

func NewFindByIDShortURLHandler(useCase *usecase.FindByIDShortURLUseCase) *FindByIDShortURLHandler {
	return &FindByIDShortURLHandler{
		useCase: useCase,
	}
}

func (h *FindByIDShortURLHandler) Handle(c *gin.Context) {

	var request dto.FindByIDShortURLRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	shortURL, err := h.useCase.Execute(request.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to find URL"})
		return
	}

	c.JSON(200, gin.H{"short_url": shortURL})
}
