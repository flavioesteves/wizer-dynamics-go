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

type ExerciseHandler struct {
	store db.MongoDBStorer
}

func NewExerciseHandler(eStore db.MongoDBStorer) *ExerciseHandler {
	return &ExerciseHandler{
		store: eStore,
	}
}

func (h *ExerciseHandler) GetAllExercises(c *gin.Context) {

	val, err := h.store.RedisClient.Get(c, "exercises").Result()
	if err == redis.Nil {
		log.Printf("Request to MongoDB")

		exercises, err := h.store.GetALlExercises(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data, _ := json.Marshal(exercises)
		h.store.RedisClient.Set(c, "exercises", string(data), 0)

		c.IndentedJSON(http.StatusOK, exercises)
	} else {
		log.Printf("Request to Redis --> exercises")
		exercises := make([]models.Exercise, 0)
		json.Unmarshal([]byte(val), &exercises)
		c.IndentedJSON(http.StatusOK, exercises)
	}
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

	h.store.RedisClient.Del(c, "exercises")

	c.JSON(http.StatusOK, exercise)
}

func (h *ExerciseHandler) UpdateExerciseById(c *gin.Context) {
	id := c.Param("id")

	exercise, err := h.store.GetExerciseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var e models.Exercise
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: Rearrange accordly to the types
	//since currently all fields are strings, "nil" is given an error
	if e.Name != "" {
		exercise.Name = e.Name
	}
	if e.Steps != "" {
		exercise.Steps = e.Steps
	}
	if e.Video != "" {
		exercise.Video = e.Video
	}
	if e.Photo != "" {
		exercise.Photo = e.Photo
	}

	exerciseUpdated, err := h.store.UpdateExerciseByID(c, exercise)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "exercises")
	c.JSON(http.StatusOK, exerciseUpdated)
}

func (h *ExerciseHandler) DeleteExerciseById(c *gin.Context) {
	id := c.Param("id")

	exercise, err := h.store.DeleteExerciseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "exercises")
	c.JSON(http.StatusOK, exercise)
}
