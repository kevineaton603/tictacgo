package main

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/kevineaton603/tictacgo/models"
	"github.com/kevineaton603/tictacgo/templates/components"
	"github.com/kevineaton603/tictacgo/templates/views"
	"github.com/labstack/echo/v4"
)

func RedirectToGame(c echo.Context) error {
	var id = uuid.New()
	return c.Redirect(301, "game/"+id.String())
}

func main() {
	e := echo.New()

	e.Static("/assets", "assets")
	e.GET("/", RedirectToGame)
	e.GET("game", RedirectToGame)
	e.GET("/game/:id", func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return RedirectToGame(c)
		}
		cells := make([]models.Player, 9)
		for i := range cells {
			cells[i] = 0
		}
		board := models.Board{Cells: cells}
		component := views.Index(models.Game{Id: id, Board: board})
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("game/:id/cell/:index", func(c echo.Context) error {
		id, idErr := uuid.Parse(c.Param("id"))
		if idErr != nil {
			return idErr
		}
		index, indexErr := strconv.Atoi(c.Param("index"))
		if indexErr != nil {
			return indexErr
		}
		component := components.Cell(index, models.X, id)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
