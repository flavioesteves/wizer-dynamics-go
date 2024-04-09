package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

type UserHandler struct {
	store db.MongoDBStorer
}

func NewUserHandler(uStore db.MongoDBStorer) *UserHandler {
	return &UserHandler{
		store: uStore,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.store.GetALlUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.store.InsertUser(c, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)

}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.store.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)

}
func (h *UserHandler) UpdateUserByID(c *gin.Context) {}
