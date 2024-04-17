package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	redisSession "github.com/gin-contrib/sessions/redis"
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

	// Redis Session
	store, _ := redisSession.NewStore(10, "tcp", redisClient.Options().Addr, "", []byte("secret"))
	router.Use(sessions.Sessions("users_api", store))

	//Stores
	exerciseStore := db.NewMongoDBStore(mDB, redisClient, "exercises")
	trainingSessionStore := db.NewMongoDBStore(mDB, redisClient, "training_sessions")
	userStore := db.NewMongoDBStore(mDB, redisClient, "users")
	authStore := db.NewMongoDBStore(mDB, redisClient, "users")
	// Handlers
	exerciseHandler := handlers.NewExerciseHandler(*exerciseStore)
	trainingSessionsHandler := handlers.NewTrainingPlanHandler(*trainingSessionStore)
	userHandler := handlers.NewUserHandler(*userStore)
	authHandler := handlers.NewAuthHandler(*authStore)
	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", handlers.GetAppStatus)
		public.POST("/signin", authHandler.SignInHandler)
		public.POST("/refresh", authHandler.RefreshHandler)
	}

	exerciseGroup := router.Group("/v1/exercises")
	trainingSessionGroup := router.Group("/v1/training-sessions")
	userGroup := router.Group("v1/users")
	v1.RegisterExerciseRoutes(exerciseGroup, *exerciseHandler)
	v1.RegisterTrainingPlanRoutes(trainingSessionGroup, *trainingSessionsHandler)
	v1.RegisterUsersRoutes(userGroup, *userHandler)

	return router
}
