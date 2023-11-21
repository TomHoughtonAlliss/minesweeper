package controllers

import (
	"context"
	//"encoding/json"
	"minesweeper/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Score struct {
	Id		primitive.ObjectID	`bson:"_id"`
	Name 	string
	Time 	int
	Date	string
}

func GetScores(c echo.Context) error {
	coll := database.Client.Database("Minesweeper").Collection("Score")

	findOptions := options.Find()
	scores, err := coll.Find(context.TODO(), bson.D{}, findOptions)
	if (err != nil) {
		panic(err)
	}

	var results []Score
	if err = scores.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, results)
}

type ScoreCreate struct {
	Name string `json:"name"`
	Time int `json:"time"`
	Date string `json:"date"`
}
func CreateScore(c echo.Context) error {

	var score ScoreCreate
	err := c.Bind(&score)
	if err != nil {
		panic("Failed to get data for score creation.")
	}

	coll := database.Client.Database("Minesweeper").Collection("Score")

	result, err := coll.InsertOne(context.TODO(), score)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}