package http

import (
	"log"
	"url/config"

	"github.com/putradwinandap/tinygo-micro/shared-lib/cache/redis"
	"github.com/putradwinandap/tinygo-micro/shared-lib/db/postgres"
	"github.com/putradwinandap/tinygo-micro/shared-lib/message_broker/rabbitmq"

	handler "url/internal/handler/shorturl"
	"url/internal/repository/db_postgres"
	usecase "url/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {

	redisClient, err := redis.NewClient(cfg.RedisURL)

	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	redisCacheCounter := redis.NewRedisCacheCounter(redisClient)

	rabbitConn, err := rabbitmq.NewRabbitMQBroker(cfg.RabbitMQURL)
	if err != nil {
		log.Fatal("RabbitMQ connection failed:", err)
	}

	rabbitPub := rabbitmq.NewRabbitMQPublisher(rabbitConn)

	db, err := postgres.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	repo := db_postgres.NewShortURLRepository(db)

	// Usecases
	saveUC := usecase.NewSaveShortURLUseCase(repo)
	findUC := usecase.NewFindByIDShortURLUseCase(repo)
	resolveUC := usecase.NewResolveShortURLUseCase(repo, rabbitPub, redisCacheCounter)

	// Handlers
	saveHandler := handler.NewSaveShortURLHandler(saveUC)
	findHandler := handler.NewFindByIDShortURLHandler(findUC)
	resolveHandler := handler.NewResolveShortURLHandler(resolveUC)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/shorturl", saveHandler.Handle)
		api.GET("/shorturl/:id", findHandler.Handle)
		api.GET("/resolve/:shortcode", resolveHandler.Handle)
	}

	return r
}
