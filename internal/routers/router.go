package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers/api/v1"
)

func SetupRouter(mDB *mongo.Database) *gin.Engine {
	router := gin.Default()

	//Stores
	exerciseStore := db.NewMongoDBStore(mDB, "exercises")
	trainingSessionStore := db.NewMongoDBStore(mDB, "training_sessions")
	userStore := db.NewMongoDBStore(mDB, "users")
	// Handlers
	exerciseHandler := handlers.NewExerciseHandler(*exerciseStore)
	trainingSessionsHandler := handlers.NewTrainingPlanHandler(*trainingSessionStore)
	userHandler := handlers.NewUserHandler(*userStore)
	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", handlers.GetAppStatus)
	}

	exerciseGroup := router.Group("/v1/exercises")
	trainingSessionGroup := router.Group("/v1/training-sessions")
	userGroup := router.Group("v1/users")
	v1.RegisterExerciseRoutes(exerciseGroup, *exerciseHandler)
	v1.RegisterTrainingPlanRoutes(trainingSessionGroup, *trainingSessionsHandler)
	v1.RegisterUsersRoutes(userGroup, *userHandler)

	return router
}
