package handler

import (
	"minesweeper/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) SaveScore(c echo.Context) (err error) {
	record := &model.Record{ID: bson.NewObjectId()}

	database := h.DB.Clone()
	defer database.Close()

	// This value will be saved as nil if there is no error
	insertError := database.DB("minesweeper").C("records").Insert(record)
	
	// If the above failed, return the relevant error
	if insertError != nil {
		return
	}

	return c.JSON(http.StatusCreated, record)
}