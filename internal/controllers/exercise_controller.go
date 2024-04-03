package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllExercises(c *gin.Context) {
	exercices := gin.H{
		"ID":    "asdaff",
		"Name":  "E1",
		"Steps": "21",
		"Video": "V1",
		"Photo": "Photo1",
	}

	c.IndentedJSON(http.StatusOK, exercices)
}
