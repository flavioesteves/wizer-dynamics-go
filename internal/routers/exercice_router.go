package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-dynamics-go/internal/controllers"
)

func RegisterExerciseRoutes(rg *gin.RouterGroup) {

	rg.GET("", controllers.GetAllExercises)
}
