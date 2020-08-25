package http

import (
	officeMysql "challenge.com/challenge/internal/office/infrastructure/persistence/mysql"
	"challenge.com/challenge/internal/search/application"
	"challenge.com/challenge/internal/shared/infrastructure"
)

// NewSearchHandlerFactory factoría para crear un Handler con sus dependencias
func NewSearchHandlerFactory(container *infrastructure.Container) *Handler {
	repository := officeMysql.NewRepository(container.Database.DB)
	byProximityService := application.NewByProximityService(repository, container.GraphHopper)

	return NewHandler(byProximityService)
}
