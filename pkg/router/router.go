package router

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/will-x86/r-place/pkg/routes/canvas"
	"github.com/will-x86/r-place/ui"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		masterKey := os.Getenv("MASTER_API_KEY")
		if masterKey == "" {
			log.Println("FATAL: MASTER_API_KEY environment variable not set.")
			http.Error(w, "Server Configuration Error", http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token != masterKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func NewRouter() *chi.Mux {
	// Create a new http router
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // Change later
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
			"PATCH",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		ExposedHeaders: []string{
			"Link",
		},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger) // Built in logger middleware
	// All /api routes go within this block
	r.Route("/api", func(r chi.Router) {
		// Create a group of *rate limited* endpoints, this will be our api endpoints :)
		r.Group(func(r chi.Router) {
			r.Use(httprate.LimitByIP(10, time.Minute)) // Limit to 10 requests per miniute within this group
			// Set a single pixel
		})
		// Move to not this lol
		r.Post("/pixels", canvas.SetCanvas)
		r.Get("/canvas", canvas.GetCanvas)
		r.Post("/pixel", canvas.GetSingle)
		r.Route("/admin", func(r chi.Router) {
			r.Use(APIKeyAuth)
			r.Delete("/section", canvas.DeleteCanvasSection)
		})
	})

	staticFileServer(r)
	return r
}
func staticFileServer(r chi.Router) {
	buildFS, err := fs.Sub(ui.StaticFiles, "build")
	if err != nil {
		log.Fatal("failed to create sub-filesystem for build directory:", err)
	}

	fileServer := http.FileServer(http.FS(buildFS))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		fsPath := strings.TrimPrefix(path, "/")
		if _, err := buildFS.Open(fsPath); err != nil {
			if os.IsNotExist(err) {
				// The file does't exist, probably client side route
				indexHTML, err := fs.ReadFile(buildFS, "index.html")
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					log.Printf("error reading index.html: %v", err)
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write(indexHTML)
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("error accessing file in embedded fs: %v", err)
			return
		}

		fileServer.ServeHTTP(w, r)
	})
}
