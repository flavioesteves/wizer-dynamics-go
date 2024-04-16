package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers/api/v1"
)

func SetupRouter(mDB *mongo.Database, redisClient *redis.Client) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	//Stores
	exerciseStore := db.NewMongoDBStore(mDB, redisClient, "exercises")
	trainingSessionStore := db.NewMongoDBStore(mDB, redisClient, "training_sessions")
	userStore := db.NewMongoDBStore(mDB, redisClient, "users")
	// Handlers
	exerciseHandler := handlers.NewExerciseHandler(*exerciseStore)
	trainingSessionsHandler := handlers.NewTrainingPlanHandler(*trainingSessionStore)
	userHandler := handlers.NewUserHandler(*userStore)
	authHandler := &handlers.AuthHandler{}
	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", handlers.GetAppStatus)
		public.POST("/signin", authHandler.SignInHandler)
	}

	exerciseGroup := router.Group("/v1/exercises")
	trainingSessionGroup := router.Group("/v1/training-sessions")
	userGroup := router.Group("v1/users")
	v1.RegisterExerciseRoutes(exerciseGroup, *exerciseHandler)
	v1.RegisterTrainingPlanRoutes(trainingSessionGroup, *trainingSessionsHandler)
	v1.RegisterUsersRoutes(userGroup, *userHandler)

	return router
}
