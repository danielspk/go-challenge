package http

import (
	"challenge.com/challenge/internal/healthcheck/database/application"
	"challenge.com/challenge/internal/healthcheck/database/infrastructure/persistence/mysql"
	"challenge.com/challenge/internal/shared/infrastructure"
)

// NewDatabaseHandlerFactory factoría para crear un Handler con sus dependencias
func NewDatabaseHandlerFactory(container *infrastructure.Container) *Handler {
	repository := mysql.NewRepository(container.Database.DB)
	pingService := application.NewPingService(repository)

	return NewHandler(pingService)
}
