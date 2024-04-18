package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	redisSession "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers/api/v1"
)

func SetupRouter(mDB *mongo.Database, redisClient *redis.Client, jwtSettings *config.JWTSettings) *gin.Engine {

	router := gin.Default()
	//router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		//AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length", "set-cookie"},
		AllowCredentials: true,
		//MaxAge:           12 * time.Hour,
	}))
	// Redis Session
	store, _ := redisSession.NewStore(10, "tcp", redisClient.Options().Addr, "", []byte(jwtSettings.Secret))
	store.Options(sessions.Options{
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		MaxAge:   jwtSettings.MaxAge,
	})

	router.Use(sessions.Sessions("users_api", store))

	//Stores
	exerciseStore := db.NewMongoDBStore(mDB, redisClient, "exercises", nil)
	trainingSessionStore := db.NewMongoDBStore(mDB, redisClient, "training_sessions", nil)
	userStore := db.NewMongoDBStore(mDB, redisClient, "users", nil)
	authStore := db.NewMongoDBStore(mDB, redisClient, "users", jwtSettings)
	// Handlers
	exerciseHandler := handlers.NewExerciseHandler(*exerciseStore)
	trainingSessionsHandler := handlers.NewTrainingPlanHandler(*trainingSessionStore)
	userHandler := handlers.NewUserHandler(*userStore)
	authHandler := handlers.NewAuthHandler(*authStore)
	// Public routes

	public := router.Group("/")
	{
		public.GET("/v1/healthcheck", handlers.GetAppStatus)
		public.POST("/v1/signin", authHandler.SignInHandler)
		public.POST("/v1/refresh", authHandler.RefreshHandler)
	}

	exerciseGroup := router.Group("/v1/exercises")
	exerciseGroup.Use(authHandler.AuthMiddleware())
	trainingSessionGroup := router.Group("/v1/training-sessions")
	userGroup := router.Group("/v1/users")

	v1.RegisterExerciseRoutes(exerciseGroup, *exerciseHandler)
	v1.RegisterTrainingPlanRoutes(trainingSessionGroup, *trainingSessionsHandler)
	v1.RegisterUsersRoutes(userGroup, *userHandler)
	return router
}
