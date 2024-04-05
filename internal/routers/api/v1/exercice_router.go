package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-dynamics-go/internal/controllers"
)

func RegisterExerciseRoutes(rg *gin.RouterGroup, eHandler controllers.ExerciseController) {

	rg.GET("", eHandler.GetAllExercises)
	rg.POST("", controllers.AddExercise)

	rg.GET("/:id", controllers.GetExerciseById)
	rg.PUT("/:id", controllers.UpdateExerciseById)
	rg.DELETE("/:id", controllers.DeleteExerciseById)
}
