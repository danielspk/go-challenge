package api

import (
	database "challenge.com/challenge/internal/healthcheck/database/infrastructure/http"
	office "challenge.com/challenge/internal/office/infrastructure/http"
	search "challenge.com/challenge/internal/search/infrastructure/http"
	"challenge.com/challenge/internal/shared/infrastructure/http/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"net/http"
)

// Handlers estructura de handlers de API
type Handlers struct {
	Office   *office.Handler
	Search   *search.Handler
	Database *database.Handler
}

// Router estructura con el router de API
type Router struct {
	http.Handler
}

// NewRouter crea un Router
func NewRouter(handlers *Handlers) *Router {
	mux := chi.NewRouter()

	mux.Use(chiMiddleware.Logger)
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(middleware.CORS)

	mux.Route("/api/v1", func(mux chi.Router) {
		handlers.Office.MakeRoutes(mux)
	})

	mux.Route("/rpc/v1", func(mux chi.Router) {
		mux.Route("/searches", func(mux chi.Router) {
			handlers.Search.MakeRoutes(mux)
		})

		mux.Route("/systems/health-checks", func(mux chi.Router) {
			handlers.Database.MakeRoutes(mux)
		})
	})

	return &Router{mux}
}
