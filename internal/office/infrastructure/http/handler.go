package http

import (
	"challenge.com/challenge/internal/office/application"
	"challenge.com/challenge/internal/shared/infrastructure/http/response"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// Handler estructura con handlers http
type Handler struct {
	findByIdService *application.FindByIdService
	createService   *application.CreateService
}

// NewHandler crea un Handler
func NewHandler(
	findByIdService *application.FindByIdService,
	createService *application.CreateService,
) *Handler {
	return &Handler{
		findByIdService: findByIdService,
		createService:   createService,
	}
}

// MakeRoutes genera las rutas para los handlers del paquete
func (h *Handler) MakeRoutes(mux chi.Router) {
	mux.Get("/offices/{id:[0-9]+}", h.show)
	mux.Post("/offices", h.create)
}

// show handler para el buscar una sucursal por ID
func (h *Handler) show(w http.ResponseWriter, r *http.Request) {
	officeID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.JSONError("the ID parameter is not valid", http.StatusBadRequest, w)
		return
	}

	cmd := &application.FindByIdCommand{ID: officeID}
	office, err := h.findByIdService.Execute(cmd)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response.JSONError("office not found", http.StatusNotFound, w)
		} else {
			response.JSONError("internal error", http.StatusInternalServerError, w)
		}
		return
	}

	response.JSON(office, http.StatusOK, w)
}

// create handler para crear una nueva sucursal
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var cmd *application.CreateCommand

	err := json.NewDecoder(r.Body).Decode(&cmd)

	if err != nil {
		response.JSONError("the body of the request is not valid", http.StatusBadRequest, w)
		return
	}

	office, messages, err := h.createService.Execute(cmd)

	if err != nil {
		if len(messages) > 0 {
			response.JSONErrors("validation with errors", messages, http.StatusBadRequest, w)

			return
		}

		response.JSONError("internal error", http.StatusInternalServerError, w)
		return
	}

	response.JSON(office, http.StatusCreated, w)
}
