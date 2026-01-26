package handler

import (
	"context"
	"net/http"
	"university/config"
	"university/database"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	DB  *database.DBConnWrapper
	Cfg *config.AppConfig
	// inject services if needed
}

func NewHealthHandler(db *database.DBConnWrapper, cfg *config.AppConfig) *HealthHandler {
	return &HealthHandler{
		DB:  db,
		Cfg: cfg,
	}
}

func (h *HealthHandler) Status(c echo.Context) error {
	// Check DB connection
	dbStatus := "ok"
	if err := h.DB.Ping(context.Background()); err != nil {
		dbStatus = "error: " + err.Error()
	}

	// Return status JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"service": "healthy",
		"db":      dbStatus,
		"version": "1.0.0",
	})
}
