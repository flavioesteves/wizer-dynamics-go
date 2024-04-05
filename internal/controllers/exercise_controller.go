package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

type ExerciseController struct {
	store db.MongoDBStorer
}

func NewExerciseController(eStore db.MongoDBStorer) *ExerciseController {
	return &ExerciseController{
		store: eStore,
	}
}

func (h *ExerciseController) GetAllExercises(c *gin.Context) {
	cursor, err := h.store.DB.Collection(h.store.Coll).Find(context.TODO(), bson.D{{}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var exercises []models.Exercise
	if err = cursor.All(context.TODO(), &exercises); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, exercises)
}

func AddExercise(c *gin.Context) {}

func GetExerciseById(c *gin.Context) {}

func UpdateExerciseById(c *gin.Context) {}

func DeleteExerciseById(c *gin.Context) {}
