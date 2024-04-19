package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/flavioesteves/wizer-dynamics-go/internal/db"
	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
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

	val, err := h.store.RedisClient.Get(c, "trainings").Result()
	if err == redis.Nil {
		log.Printf("Request to MongoDB")

		trainingSessions, err := h.store.GetALlTrainings(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data, _ := json.Marshal(trainingSessions)
		h.store.RedisClient.Set(c, "trainings", string(data), 0)
		c.IndentedJSON(http.StatusOK, trainingSessions)
	} else {
		log.Printf("Request to Redis --> trainings")
		trainingSessions := make([]models.TrainingSession, 0)
		json.Unmarshal([]byte(val), &trainingSessions)
		c.IndentedJSON(http.StatusOK, trainingSessions)
	}
}

func (h *TrainingPlanHandler) GetTrainingById(c *gin.Context) {
	id := c.Param("id")
	trainingPlan, err := h.store.GetTrainigByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, trainingPlan)
}

func (h *TrainingPlanHandler) AddTraining(c *gin.Context) {
	var tp models.TrainingSession

	if err := c.ShouldBindJSON(&tp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	trainingPlan, err := h.store.InsertTraining(c, &tp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "trainings")
	c.JSON(http.StatusOK, trainingPlan)
}

func (h *TrainingPlanHandler) UpdateTrainingById(c *gin.Context) {
	id := c.Param("id")

	trainingSession, err := h.store.GetTrainigByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var ts models.TrainingSession
	if err := c.ShouldBindJSON(&ts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Todo: Rearrange accordly to the types
	//since currently all fields are strings, "nil" is given an error
	if ts.Day != "" {
		trainingSession.Day = ts.Day
	}
	if ts.Theme != "" {
		trainingSession.Theme = ts.Theme
	}
	if ts.ScheduleDays != "" {
		trainingSession.ScheduleDays = ts.ScheduleDays
	}
	if ts.EstimatedTime != "" {
		trainingSession.EstimatedTime = ts.EstimatedTime
	}

	tSessionUpdated, err := h.store.UpdateTrainingByID(c, trainingSession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "trainings")
	c.JSON(http.StatusOK, tSessionUpdated)

}

func (h *TrainingPlanHandler) DeleteTrainingByID(c *gin.Context) {
	id := c.Param("id")

	trainingSession, err := h.store.DeleteTrainingByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "trainings")
	c.JSON(http.StatusOK, trainingSession)
}
