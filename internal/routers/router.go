package routers

import (
	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/controllers"
	"github.com/flavioesteves/wizer-dynamics-go/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(dbClient mongo.Client, dbConfig *config.DatabaseSettings) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", dbClient)
		c.Set("dbConfig", dbConfig)
	})

	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", controllers.GetAppStatus)
	}

	exerciseGroup := router.Group("/v1/exercise")
	trainingPlanGroup := router.Group("/v1/training-plan")
	v1.RegisterExerciseRoutes(exerciseGroup)
	v1.RegisterTrainingPlanRoutes(trainingPlanGroup)

	return router
}
