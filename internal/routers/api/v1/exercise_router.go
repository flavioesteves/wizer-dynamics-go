package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
)

func RegisterExerciseRoutes(rg *gin.RouterGroup, h handlers.ExerciseHandler) {

	rg.GET("", h.GetAllExercises)
	rg.POST("", h.AddExercise)

	rg.GET("/:id", h.GetExerciseById)
	rg.PUT("/:id", h.UpdateExerciseById)
	rg.DELETE("/:id", h.DeleteExerciseById)
}
