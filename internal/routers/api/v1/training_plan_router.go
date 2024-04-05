package v1

import (
	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTrainingPlanRoutes(rg *gin.RouterGroup, h handlers.TrainingPlanHandler) {
	rg.GET("", h.GetALlTrainings)
	rg.POST("", h.AddTraining)

	rg.GET("/:id", h.GetTrainingById)
	rg.PUT("/:id", h.UpdateTrainingById)
	rg.DELETE("/:id", h.DeleteTrainingByID)
}
