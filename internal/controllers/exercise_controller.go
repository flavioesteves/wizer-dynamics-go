package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/flavioesteves/wizer-dynamics-go/configs"
	"github.com/flavioesteves/wizer-dynamics-go/internal/models"
)

func GetAllExercises(c *gin.Context) {

	db := c.MustGet("db").(mongo.Client)
	dbConfig := c.MustGet("dbConfig").(*config.DatabaseSettings)

	cursor, err := db.Database(dbConfig.DatabaseName).Collection("exercises").Find(context.TODO(), bson.D{{}})

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
