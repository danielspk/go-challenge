package api

import (
	"fmt"
	"net/http"
	"time"
)

//Server estructura de servidor HTTP
type Server struct {
	*http.Server
}

// NewServer crea un Server configurado
func NewServer(port uint16, handler http.Handler) *Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{server}
}
