package http

import (
	"challenge.com/challenge/internal/search/application"
	"challenge.com/challenge/internal/shared/infrastructure/http/response"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// Handler estructura con handlers http
type Handler struct {
	byProximityService *application.ByProximityService
}

// NewHandler crea un Handler
func NewHandler(byProximityService *application.ByProximityService) *Handler {
	return &Handler{
		byProximityService: byProximityService,
	}
}

// MakeRoutes genera las rutas para los handlers del paquete
func (h *Handler) MakeRoutes(mux chi.Router) {
	mux.Get("/officeByProximity", h.findByProximity)
}

// findByProximity handler para buscar la sucursal más cercana
func (h *Handler) findByProximity(w http.ResponseWriter, r *http.Request) {
	queryLatitude := r.URL.Query().Get("latitude")
	queryLongitude := r.URL.Query().Get("longitude")

	if queryLatitude == "" {
		response.JSONError("empty latitude param", http.StatusBadRequest, w)
		return
	}

	if queryLongitude == "" {
		response.JSONError("empty longitude param", http.StatusBadRequest, w)
		return
	}

	latitude, err := strconv.ParseFloat(queryLatitude, 32)

	if err != nil {
		response.JSONError("invalid latitude format", http.StatusBadRequest, w)
		return
	}

	longitude, err := strconv.ParseFloat(queryLongitude, 32)

	if err != nil {
		response.JSONError("invalid longitude format", http.StatusBadRequest, w)
		return
	}

	cmd := &application.ByProximityCommand{
		Latitude:  float32(latitude),
		Longitude: float32(longitude),
	}

	proximityResponse, err := h.byProximityService.Execute(cmd)

	if err != nil {
		response.JSONError("no office found", http.StatusBadRequest, w)
		return
	}

	response.JSON(proximityResponse, http.StatusOK, w)
}
