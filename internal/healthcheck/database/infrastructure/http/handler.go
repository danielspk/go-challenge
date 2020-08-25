package http

import (
	"challenge.com/challenge/internal/healthcheck/database/application"
	"challenge.com/challenge/internal/shared/infrastructure/http/response"
	"github.com/go-chi/chi"
	"net/http"
)

// Handler estructura con handlers http
type Handler struct {
	pingService *application.PingService
}

// NewHandler crea un Handler
func NewHandler(pingService *application.PingService) *Handler {
	return &Handler{
		pingService: pingService,
	}
}

// MakeRoutes genera las rutas para los handlers del paquete
func (h *Handler) MakeRoutes(mux chi.Router) {
	mux.Get("/database/ping", h.ping)
}

// ping handler para el health check de la base de datos
func (h *Handler) ping(w http.ResponseWriter, _ *http.Request) {
	err := h.pingService.Execute()

	if err != nil {
		response.PlainText("ERROR", http.StatusInternalServerError, w)
		return
	}

	response.PlainText("PONG", http.StatusOK, w)
}
