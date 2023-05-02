package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"ChoreChamps/model"
)

type KidHandler struct {
	db *gorm.DB
}

func NewKidHandler(db *gorm.DB) *KidHandler {
	return &KidHandler{db: db}
}

func (h *KidHandler) Create(c echo.Context) error {
	// Implement the create kid handler
	return c.JSON(http.StatusNotImplemented, nil)
}

func (h *KidHandler) Get(c echo.Context) error {
	// Implement the get kid handler
	return c.JSON(http.StatusNotImplemented, nil)
}

func (h *KidHandler) GetAll(c echo.Context) error {
	// Implement the get all kids handler
	return c.JSON(http.StatusNotImplemented, nil)
}

func (h *KidHandler) Update(c echo.Context) error {
	// Implement the update kid handler
	return c.JSON(http.StatusNotImplemented, nil)
}

func (h *KidHandler) Delete(c echo.Context) error {
	// Implement the delete kid handler
	return c.JSON(http.StatusNotImplemented, nil)
}