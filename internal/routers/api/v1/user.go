package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-dynamics-go/internal/handlers"
)

func RegisterUsersRoutes(rg *gin.RouterGroup, h handlers.UserHandler) {

	rg.GET("", h.GetAllUsers)
	rg.POST("", h.AddUser)

	rg.GET("/:id", h.GetUserByID)
	rg.PUT("/:id", h.UpdateUserByID)
}
