package http

/*
import (
	"url/internal/config"
	"url/internal/db"
	handler "url/internal/handler/shorturl"
	"url/internal/repository/db_postgres"
	usecase "url/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
)

func saveShortURLHandler(c *gin.Context) {

	loadConf := config.LoadConfig()
	PostgresDB, err := db.NewPostgresDB(loadConf.DatabaseURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}

	repo := db_postgres.NewShortURLRepository(PostgresDB)
	useCase := usecase.NewSaveShortURLUseCase(repo)
	handler := handler.NewSaveShortURLHandler(useCase)

	handler.Handle(c)
}

func getShortURLByIDHandler(c *gin.Context) {
	loadConf := config.LoadConfig()
	PostgresDB, err := db.NewPostgresDB(loadConf.DatabaseURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	repo := db_postgres.NewShortURLRepository(PostgresDB)
	useCase := usecase.NewFindByIDShortURLUseCase(repo)
	handler := handler.NewFindByIDShortURLHandler(useCase)
	handler.Handle(c)
}

func resolveShortURLHandler(c *gin.Context) {
	loadConf := config.LoadConfig()
	PostgresDB, err := db.NewPostgresDB(loadConf.DatabaseURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	repo := db_postgres.NewShortURLRepository(PostgresDB)
	useCase := usecase.NewResolveShortURLUseCase(repo)
	handler := handler.NewResolveShortURLHandler(useCase)
	handler.Handle(c)
}
*/
