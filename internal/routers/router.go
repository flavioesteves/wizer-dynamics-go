package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/internal/controllers"
	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers/api/v1"
)

func SetupRouter(mDB *mongo.Database) *gin.Engine {
	router := gin.Default()

	exerciseStore := db.NewMongoDBStore(mDB, "exercises")
	exerciseController := controllers.NewExerciseController(*exerciseStore)

	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", controllers.GetAppStatus)
	}

	exerciseGroup := router.Group("/v1/exercise")
	trainingPlanGroup := router.Group("/v1/training-plan")
	v1.RegisterExerciseRoutes(exerciseGroup, *exerciseController)
	v1.RegisterTrainingPlanRoutes(trainingPlanGroup)

	return router
}
