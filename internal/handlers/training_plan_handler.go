package handlers

import (
	"context"
	"net/http"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type TrainingPlanHandler struct {
	store db.MongoDBStorer
}

func NewTrainingPlanHandler(tpStore db.MongoDBStorer) *TrainingPlanHandler {
	return &TrainingPlanHandler{
		store: tpStore,
	}
}

func (h *TrainingPlanHandler) GetALlTrainings(c *gin.Context) {
	cursor, err := h.store.DB.Collection(h.store.Coll).Find(context.TODO(), bson.D{{}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var trainingPlans []models.TrainingPlan
	if err = cursor.All(context.TODO(), &trainingPlans); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, trainingPlans)

}

func (h *TrainingPlanHandler) AddTraining(c *gin.Context) {}

func (h *TrainingPlanHandler) GetTrainingById(c *gin.Context) {}

func (h *TrainingPlanHandler) UpdateTrainingById(c *gin.Context) {}

func (h *TrainingPlanHandler) DeleteTrainingByID(c *gin.Context) {}
