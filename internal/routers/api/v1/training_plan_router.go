package v1

import (
	"github.com/flavioesteves/wizer-dynamics-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTrainingPlanRoutes(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetALlTrainings)
	rg.POST("", controllers.AddTraining)

	rg.GET("/:id", controllers.GetTrainingById)
	rg.PUT("/:id", controllers.UpdateTrainingById)
	rg.DELETE("/:id", controllers.DeleteExerciseById)
}
