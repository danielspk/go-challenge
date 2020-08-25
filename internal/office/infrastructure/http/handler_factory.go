package http

import (
	"challenge.com/challenge/internal/office/application"
	"challenge.com/challenge/internal/office/infrastructure/persistence/mysql"
	"challenge.com/challenge/internal/shared/infrastructure"
)

// NewOfficeHandlerFactory factoría para crear un Handler con sus dependencias
func NewOfficeHandlerFactory(container *infrastructure.Container) *Handler {
	repository := mysql.NewRepository(container.Database.DB)
	findByIdService := application.NewFindByIdService(repository)
	createService := application.NewCreateService(repository)

	return NewHandler(findByIdService, createService)
}
