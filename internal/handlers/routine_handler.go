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

type RoutineHandler struct {
	store db.MongoDBStorer
}

func NewRoutineHandler(rStore db.MongoDBStorer) *RoutineHandler {
	return &RoutineHandler{
		store: rStore,
	}
}

func (h *RoutineHandler) GetALlRoutines(c *gin.Context) {

	val, err := h.store.RedisClient.Get(c, "routines").Result()
	if err == redis.Nil {
		log.Printf("Request to MongoDB")

		routines, err := h.store.GetALlRoutines(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data, _ := json.Marshal(routines)
		h.store.RedisClient.Set(c, "routines", string(data), 0)
		c.IndentedJSON(http.StatusOK, routines)
	} else {
		log.Printf("Request to Redis --> trainings")
		routines := make([]models.Routine, 0)
		json.Unmarshal([]byte(val), &routines)
		c.IndentedJSON(http.StatusOK, routines)
	}
}

func (h *RoutineHandler) GetRoutineById(c *gin.Context) {
	id := c.Param("id")
	routine, err := h.store.GetRoutineByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, routine)
}

func (h *RoutineHandler) AddRoutine(c *gin.Context) {
	var r models.Routine

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	routine, err := h.store.InsertRoutine(c, &r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "routines")
	c.JSON(http.StatusOK, routine)
}

func (h *RoutineHandler) UpdateRoutineById(c *gin.Context) {
	id := c.Param("id")

	routine, err := h.store.GetRoutineByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var r models.Routine
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Todo: Rearrange accordly to the types
	//since currently all fields are strings, "nil" is given an error
	if r.Day != "" {
		routine.Day = r.Day
	}
	if r.Theme != "" {
		routine.Theme = r.Theme
	}
	if r.ScheduleDays != "" {
		routine.ScheduleDays = r.ScheduleDays
	}
	if r.EstimatedTime != "" {
		routine.EstimatedTime = r.EstimatedTime
	}

	rUpdated, err := h.store.UpdateRoutineByID(c, routine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "routines")
	c.JSON(http.StatusOK, rUpdated)

}

func (h *RoutineHandler) DeleteRoutineByID(c *gin.Context) {
	id := c.Param("id")

	routine, err := h.store.DeleteRoutineByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.store.RedisClient.Del(c, "trainings")
	c.JSON(http.StatusOK, routine)
}
