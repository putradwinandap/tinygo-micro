package shorturl

import (
	"url/internal/dto"
	usecase "url/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
)

type ResolveShortURLHandler struct {
	useCase *usecase.ResolveShortURLUseCase
}

func NewResolveShortURLHandler(useCase *usecase.ResolveShortURLUseCase) *ResolveShortURLHandler {
	return &ResolveShortURLHandler{
		useCase: useCase,
	}
}

func (h *ResolveShortURLHandler) Handle(c *gin.Context) {

	var request dto.ResolveShortURLRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	shortURL, err := h.useCase.Execute(request.ShortCode)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to find URL"})
		return
	}

	c.JSON(200, gin.H{"short_url": shortURL})
}
