package handlers

import (
	"net/http"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/gin-gonic/gin"
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
	trainingPlans, err := h.store.GetALlExercises(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, trainingPlans)
}

func (h *TrainingPlanHandler) GetTrainingById(c *gin.Context) {
	id := c.Param("id")
	trainingPlan, err := h.store.GetTrainigByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, trainingPlan)
}

func (h *TrainingPlanHandler) AddTraining(c *gin.Context) {}

func (h *TrainingPlanHandler) UpdateTrainingById(c *gin.Context) {}

func (h *TrainingPlanHandler) DeleteTrainingByID(c *gin.Context) {}
