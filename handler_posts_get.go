package main

import (
	"context"
	"log"
	"net/http"
    "strconv"

	_ "github.com/lib/pq"
	"github.com/skovranek/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user User) {
    limit := int32(10)
	limitStr := r.URL.Query().Get("author_id")
	if limitStr != "" {
        limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("Error converting limit str to int: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Unable to get limit query parameter")
			return
		}
        limit = int32(limitInt)
	}

	ctx := context.Background()

    dbPosts, err := cfg.DB.GetPostsByUser(ctx, database.GetPostsByUserParams{
        UserID: user.ID,
        Limit: limit,
    })
	if err != nil {
		log.Printf("Error: handlerPostsGet: cfg.DB.GetPostsByUser: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve posts from database")
		return
	}

	posts := []Post{}
	for _, dbPost := range dbPosts {
		post := databasePostToPost(dbPost)
		posts = append(posts, post)
    }

	respondWithJSON(w, http.StatusOK, posts)
}
