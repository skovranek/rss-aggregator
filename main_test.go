package main

import (
    "encoding/json"
    "io"
	"net/http"
    "testing"
    "time"
)

func TestMain(t *testing.T) {
    go main()

    time.Sleep(time.Second)
	t.Run("Test server readiness", func(t *testing.T) {
        res, err := http.Get("http://localhost:8080/v1/healthz")
        if err != nil {
            t.Errorf("Error: TestMain: http.Get(\"localhost.8080/healthz\"): %v", err)
            return
        }
        defer res.Body.Close()
        resBody, err := io.ReadAll(res.Body)
        if err != nil {
            t.Errorf("Error: TestMain: io.ReadAll(resp.Body): %v", err)
            return
        }
        if res.StatusCode != http.StatusOK {
            t.Errorf("Unexpected: TestMain: %v", res.StatusCode)
            return
        }

        if res.Header.Get("Content-Type") != "application/json" {
            t.Errorf("Unexpected: TestMain: %v", res.Header.Get("Content-Type"))
            return
        }

        output := struct {
		    Status string `json:"status"`
	    }{}

        err = json.Unmarshal(resBody, &output)
        if err != nil {
            t.Errorf("Error: TestMain: json.Unmarshal(resBody, &output): %v", err)
            return
        }

        if output.Status != "ok" {
            t.Errorf("Unexpected: TestMain: output.Status: %v", output.Status)
            return
        }
	})
}
