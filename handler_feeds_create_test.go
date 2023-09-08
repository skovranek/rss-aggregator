package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "time"

    "github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/skovranek/rss_aggregator/internal/database"
)

func TestHandlerFeedsCreate(t *testing.T) {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(`Error: main.go: godotenv.Load(): cannot load ".env" file`)
    }

    dbURL := os.Getenv("CONN")
    dbQueries, db := database.InitDB(dbURL)
    defer db.Close()

    cfg := apiConfig{
        DB:    dbQueries,
    }

    feedParams := FeedParams{
        Name: fmt.Sprintf("FEED_NAME: test_handler_feeds_create: %s", time.Now().String()),
        URL: fmt.Sprintf("FEED_URL: test_handler_feeds_create: %s", time.Now().String()),
    }

    marshalled, err := json.Marshal(feedParams)
    if err != nil {
        t.Errorf("Error: TestHandlerFeedsCreate: json.Marshal(feedParams): %v", err)
    }
    reqBody := bytes.NewReader(marshalled)

    req := httptest.NewRequest(http.MethodGet, "/", reqBody)
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()

    userID := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")
    user := User{
        ID: userID,
    }

    cfg.handlerFeedsCreate(w, req, user)

    res := w.Result()
    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        t.Errorf("Error: TestHandlerFeedsCreate: io.ReadAll(res.Body): %v", err)
        return
    }

    if res.StatusCode == http.StatusInternalServerError {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", res.StatusCode)
        return
    }

    if res.Header.Get("Content-Type") != "application/json" {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", res.Header.Get("Content-Type"))
        return
    }

    output := struct {
        Feed   Feed                `json:"feed"`
        Follow database.FeedFollow `json:"feed_follow"`
    }{}

    err = json.Unmarshal(resBody, &output)
    if err != nil {
        t.Errorf("Error: TestHandlerFeedsCreate: %v", err)
        return
    }

    if output.Feed.UserID != userID {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", output.Feed.UserID)
        return
    }

    if output.Feed.Name != feedParams.Name {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", output.Feed.Name)
        return
    }

    if output.Feed.Url != feedParams.URL {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", output.Feed.Url)
        return
    }

    if output.Follow.FeedID != output.Feed.ID {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", output.Follow.FeedID)
        return
    }

    if output.Follow.UserID != output.Feed.UserID {
        t.Errorf("Unexpected: TestHandlerFeedsCreate: %v", output.Follow.UserID)
        return
    }
}
