package api

import (
	database "challenge.com/challenge/internal/healthcheck/database/infrastructure/http"
	office "challenge.com/challenge/internal/office/infrastructure/http"
	search "challenge.com/challenge/internal/search/infrastructure/http"
	"challenge.com/challenge/internal/shared/infrastructure"
	"log"
)

// Factory factoría para crear el web server de la API
func Factory(port uint16, container *infrastructure.Container) *Server {
	apiHandlers := &Handlers{
		Office:   office.NewOfficeHandlerFactory(container),
		Search:   search.NewSearchHandlerFactory(container),
		Database: database.NewDatabaseHandlerFactory(container),
	}

	router := NewRouter(apiHandlers)
	server := NewServer(port, router)

	log.Println("go challenge API is running")

	return server
}
