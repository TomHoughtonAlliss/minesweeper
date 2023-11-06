package handler

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo/v4"
	"github.com/jamescape/goHackday-minesweeper/model"
)

func (h *Handler) SaveScore(c echo.Context) (error Error) {
	record := &model.Record{ID: bson.NewObjectId()}
}