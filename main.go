package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    "github.com/go-chi/chi"
    "github.com/go-chi/cors"
    "github.com/joho/godotenv"

	"github.com/skovranek/rss_aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
    Limit int32
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(`Error: main.go: cannot load ".env" file`)
    }
    port := os.Getenv("PORT")
    dbURL := os.Getenv("CONN")
    db, err := sql.Open("postgres", dbURL)
    dbQueries := database.New(db)
    cfg := apiConfig{
        DB: dbQueries,
        Limit: int32(10),
    }
    go cfg.workerFetchFeeds()

    r := chi.NewRouter()
    r.Use(cors.Handler(cors.Options{
        // AllowedOrigins:   []string{"https://foo.com"}, 
        // Use this to allow specific origin hosts
        AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
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
    v1router.Get("/feed_follows/", cfg.middlewareAuth(cfg.handlerFollowsGet))
    r.Mount("/v1", middlewareLog(v1router))
	
    srv := &http.Server{
        Handler: r,
        Addr:    "localhost:" + port,
    }

    log.Printf("Starting server on 'localhost:%s'", port)
    log.Fatal(srv.ListenAndServe())
}

