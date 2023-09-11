package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (cfg *apiConfig) createServer() *http.Server {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()

	v1router.Get("/readiness", handlerReadiness)
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", handlerErr)

	v1router.Post("/users", cfg.handlerUsersCreate)
	v1router.Get("/users", cfg.middlewareAuth(cfg.handlerUsersGet))

	v1router.Post("/feeds", cfg.middlewareAuth(cfg.handlerFeedsCreate))
	v1router.Get("/feeds", cfg.handlerFeedsGet)

	v1router.Post("/feed_follows", cfg.middlewareAuth(cfg.handlerFollowsCreate))
	v1router.Delete("/feed_follows/{feedFollowID}", cfg.handlerFollowsDelete)
	v1router.Get("/feed_follows", cfg.middlewareAuth(cfg.handlerFollowsGetByUser))

	v1router.Get("/posts", cfg.middlewareAuth(cfg.handlerPostsGet))

	r.Mount("/v1", middlewareLog(v1router))

	return &http.Server{
		Handler:           r,
		Addr:              "localhost:" + cfg.port,
		ReadHeaderTimeout: 10 * time.Second,
	}
}
