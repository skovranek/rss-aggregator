package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerErr(t *testing.T) {
	t.Run("TestHandlerErr", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handlerErr(w, req)

		res := w.Result()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error: TestHandlerErr: io.ReadAll(res.Body): %v", err)
			return
		}

		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Unexpected: TestHandlerErr: %v", res.StatusCode)
			return
		}

		if res.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Unexpected: TestHandlerErr: %v", res.Header.Get("Content-Type"))
			return
		}

		output := struct {
			Error string `json:"error"`
		}{}

		err = json.Unmarshal(body, &output)
		if err != nil {
			t.Errorf("Error: TestHandlerErr: %v", err)
			return
		}

		if output.Error != "Internal Server Error" {
			t.Errorf("Unexpected: TestHandlerErr: %v", output)
			return
		}
	})
}
