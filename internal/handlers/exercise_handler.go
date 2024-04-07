package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

type ExerciseHandler struct {
	store db.MongoDBStorer
}

func NewExerciseHandler(eStore db.MongoDBStorer) *ExerciseHandler {
	return &ExerciseHandler{
		store: eStore,
	}
}

func (h *ExerciseHandler) GetAllExercises(c *gin.Context) {
	exercises, err := h.store.GetALlExercises(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) GetExerciseById(c *gin.Context) {
	id := c.Param("id")

	exercise, err := h.store.GetExerciseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, exercise)
}

func (h *ExerciseHandler) AddExercise(c *gin.Context) {
	var e models.Exercise

	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := h.store.InsertExercise(c, &e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, exercise)
}

func (h *ExerciseHandler) UpdateExerciseById(c *gin.Context) {}

func (h *ExerciseHandler) DeleteExerciseById(c *gin.Context) {
	id := c.Param("id")

	exercise, err := h.store.DeleteExerciseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, exercise)
}
