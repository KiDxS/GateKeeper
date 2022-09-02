package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Application) Routes() (handler http.Handler) {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers

	}))
	// Api version 1
	router.Route("/api/v1", func(r chi.Router) {

		// User authentication
		// /api/v1/user/
		r.Route("/user", func(r chi.Router) {
			// /api/v1/user/login
			r.Post("/login", handleLogin)
			// /api/v1/user/logout
			r.Get("/logout", handleLogout)
			// Requires authentication
			r.Group(func(r chi.Router) {
				r.Use(authMiddleware)
				r.Post("/change-password", handleChangePassword)
			})

		})
		// /api/v1/protected
		r.Route("/protected", func(r chi.Router) {
			r.Use(authMiddleware)
			r.Get("/", handleIndex)
		})
	})

	return router
}
