package handler

import (
	"minesweeper/app/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) SaveScore(c echo.Context) (err error) {
	record := &model.Record{ID: bson.NewObjectId()}
}