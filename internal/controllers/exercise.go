package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetALLExercises(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, exercises)
}
