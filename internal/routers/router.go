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
	trainingPlansStore := db.NewMongoDBStore(mDB, "training_sessions")

	// Handlers
	exerciseHandler := handlers.NewExerciseHandler(*exerciseStore)
	trainingPlansHandler := handlers.NewTrainingPlanHandler(*trainingPlansStore)
	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", handlers.GetAppStatus)
	}

	exerciseGroup := router.Group("/v1/exercises")
	trainingPlanGroup := router.Group("/v1/training-sessions")
	v1.RegisterExerciseRoutes(exerciseGroup, *exerciseHandler)
	v1.RegisterTrainingPlanRoutes(trainingPlanGroup, *trainingPlansHandler)

	return router
}
