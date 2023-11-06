package handler

import (
	"minesweeper/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

type CreateRecordRequest struct {
	Name string `json:"name"`
	Score  string `json:"score"`
}

func (h *Handler) SaveRecords(c echo.Context) error {

	var request CreateRecordRequest
	err := c.Bind(&request)
	if err != nil {
		return err
	}
	record := &model.Record{ID: bson.NewObjectId(), Name: request.Name, Score: request.Score}

	database := h.DB.Clone()
	defer database.Close()

	// This value will be saved as nil if there is no error
	insertError := database.DB("minesweeper").C("records").Insert(record)
	
	// If the above failed, return the relevant error
	if insertError != nil {
		return insertError
	}

	return c.JSON(http.StatusCreated, record)
}

func (h *Handler) GetRecords(c echo.Context) error {
	database := h.DB.Clone()

	records := []*model.Record{}
	err := database.DB("minesweeper").C("records").Find(bson.M{}).All(&records)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, records)
}