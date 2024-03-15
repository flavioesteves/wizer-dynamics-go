package routes

import (
	//"encoding/json"
	"net/http"

	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
	"github.com/gin-gonic/gin"
)

func RegisterExerciseRoutes(rg *gin.RouterGroup) {
	rg.GET("", getAllExercises)
	//rg.POST("", addExercise)
	rg.GET("/:id", getExercisebyID)
	//rg.PUT("/:id", updateExerciseByID)
	//rg.DELETE(":id", deleteExerciseByID)
}

// TODO delete
var exercises = []models.Exercise{
	{ID: "afasfafs1", Name: "Exe1", Steps: "21", Video: "Video1", Photo: "Photo1"},
}

func getAllExercises(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, exercises)
}
func addExercise() {}
func getExercisebyID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range exercises {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "exercise not found"})
	}
}
func updateExerciseByID() {}
func deleteExerciseByID() {}
