package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/will-x86/r-place/pkg/routes/canvas"
)

func NewRouter() *chi.Mux {
	// Create a new http router
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Built in logger middleware
	// All /api routes go within this block
	r.Route("/api", func(r chi.Router) {
		// Create a group of *rate limited* endpoints, this will be our api endpoints :)
		r.Group(func(r chi.Router) {
			r.Use(httprate.LimitByIP(10, time.Minute)) // Limit to 10 requests per miniute within this group
			// Set a single pixel
			r.Post("/pixels", canvas.SetCanvas)
		})
		r.Get("/canvas", canvas.GetCanvas)
		r.Post("/pixel", canvas.GetSingle)
	})

	return r
}
