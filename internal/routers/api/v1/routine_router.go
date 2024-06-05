package v1

import (
	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutinesRoutes(rg *gin.RouterGroup, h handlers.RoutineHandler) {
	rg.GET("", h.GetALlRoutines)
	rg.POST("", h.AddRoutine)

	rg.GET("/:id", h.GetRoutineById)
	rg.PUT("/:id", h.UpdateRoutineById)
	rg.DELETE("/:id", h.DeleteRoutineByID)
}
