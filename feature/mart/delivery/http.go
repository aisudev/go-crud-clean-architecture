package delivery

import (
	"crud/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.MartUsecase
}

func NewHandler(e *echo.Group, u domain.MartUsecase) *Handler {
	h := Handler{usecase: u}

	e.POST("/mart", h.CreateMart)
	e.GET("/mart/:id", h.GetMart)
	e.GET("/mart", h.GetAllMart)
	e.PUT("/mart/:id", h.UpdateMart)
	e.DELETE("/mart/:id", h.DeleteMart)

	return &h
}

func (h *Handler) CreateMart(c echo.Context) error {

	mart := new(domain.Mart)

	if err := c.Bind(&mart); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false})
	}

	if err := h.usecase.CreateMart(mart); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})

}

func (h *Handler) GetMart(c echo.Context) error {
	param, _ := strconv.ParseUint(c.Param("id"), 10, 16)
	id := uint16(param)

	var mart *domain.Mart
	var err error
	if mart, err = h.usecase.GetMart(id); err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false})
	}

	return c.JSON(http.StatusOK, mart)
}

func (h *Handler) GetAllMart(c echo.Context) error {
	var mart []domain.Mart
	var err error

	if mart, err = h.usecase.GetAllMart(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, mart)
}

func (h *Handler) UpdateMart(c echo.Context) error {

	var mart domain.Mart

	param, _ := strconv.ParseUint(c.Param("id"), 10, 16)
	id := uint16(param)

	if err := c.Bind(&mart); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.usecase.UpdateMart(id, mart); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})

}

func (h *Handler) DeleteMart(c echo.Context) error {

	param, _ := strconv.ParseUint(c.Param("id"), 10, 16)
	id := uint16(param)

	if err := h.usecase.DeleteMart(id); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
