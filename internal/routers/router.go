package routers

import (
	"github.com/flavioesteves/wizer-dynamics-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/v1")
	{
		public.GET("/healthcheck", controllers.GetAppStatus)
	}

	exerciseGroup := router.Group("/v1/exercise")
	RegisterExerciseRoutes(exerciseGroup)

	return router
}
