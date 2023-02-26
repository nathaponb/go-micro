package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

/*
 * Routes map api endpoints to handlers
 */

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	/*
	 * config Cors
	 */
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)

	// grpc specific route
	mux.Post("/log-grpc", app.LogItemViaGRPC)

	mux.Post("/handle", app.HandleSubmission)

	return mux
}
