package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kevineaton603/tictacgo/models"
	"github.com/kevineaton603/tictacgo/templates/components"
	"github.com/kevineaton603/tictacgo/templates/views"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GameHandler struct {
	router *echo.Echo
	db     *gorm.DB
}

func NewGameHandler(router *echo.Echo, db *gorm.DB) *GameHandler {
	return &GameHandler{
		router: router,
		db:     db,
	}
}

func (h *GameHandler) Mount() {
	router := h.router.Group("/game")
	router.GET("/", h.newGame)
	router.GET("/:id", h.getGame)
	router.POST("/:id/cell/:index", h.updateCell)

}

func (h *GameHandler) newGame(c echo.Context) error {
	c.Logger().Print("In New Game Route")
	game := models.NewGame()
	result := h.db.Create(game)
	if result.Error != nil {
		c.Logger().Errorf(result.Error.Error())
		return c.String(500, "Failed to create model")
	}
	return c.Redirect(302, game.Id.String())
}

func (h *GameHandler) getGame(c echo.Context) error {
	c.Logger().Print("In Get Game Route")
	id, idErr := uuid.Parse(c.Param("id"))
	if idErr != nil {
		return h.newGame(c)
	}
	game := new(models.Game)
	result := h.db.First(game, "id = ?", id.String())
	if result.Error != nil {
		c.Logger().Errorf(result.Error.Error())
		return c.String(500, "Something went wrong")
	}

	component := views.Index(game)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *GameHandler) updateCell(c echo.Context) error {
	c.Logger().Print("In UpdateCell Route")
	id, idErr := uuid.Parse(c.Param("id"))
	if idErr != nil {
		return idErr
	}
	index, indexErr := strconv.Atoi(c.Param("index"))
	if indexErr != nil {
		return indexErr
	}
	entity := new(models.Game)
	result := h.db.First(entity, "id = ?", id.String())
	if result.Error != nil {
		c.Logger().Errorf(result.Error.Error())
		return c.String(http.StatusInternalServerError, "Failed to Find")
	}
	entity.UpdateCell(index)
	entity.FindWinner()
	entity.SwitchPlayer()
	update := h.db.Save(entity)
	if update.Error != nil {
		c.Logger().Errorf(update.Error.Error())
		return c.String(http.StatusInternalServerError, "Failed to Update")
	}
	component := components.Board(entity)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
