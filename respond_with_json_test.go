package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http/httptest"
    "testing"
)

func TestRespondWithJSON(t *testing.T) {
    type Payload struct {
        Payload any `json:"payload"`
    }

	tests := []struct {
        payload Payload
        statusCode int
	}{
        {}, // zero values
        {
            payload: Payload{
                Payload: "this is a string",
            },
            statusCode: 200,
        },
	}

	for i, test := range tests {
        t.Run(fmt.Sprintf("TestRespondWithJSON Case #%v:", i), func(t *testing.T) {
            //req := httptest.NewRequest(http.MethodGet, "/", nil)
            w := httptest.NewRecorder()

            respondWithJSON(w, test.statusCode, test.payload)

            res := w.Result()
            body, err := io.ReadAll(res.Body)
            if err != nil {
                t.Errorf("Error: TestRespondWithJSON: io.ReadAll(res.Body): %v", err)
                return
            }

            if res.StatusCode != test.statusCode {
                if res.StatusCode != 409 && test.statusCode > 99 && test.statusCode < 600 {
                    t.Errorf("Unexpected: TestRespondWithJSON Case #%v: %v", i, res.StatusCode)
                    return
                }
			}

			if res.Header.Get("Content-Type") != "application/json" {
                t.Errorf("Unexpected: TestRespondWithJSON Case #%v: %v", i, res.Header.Get("Content-Type"))
				return
            }

            output := Payload{}
			err = json.Unmarshal(body, &output)
			if err != nil {
                t.Errorf("Error: TestRespondWithJSON Case #%v: %v", i, err)
				return
			}

			if output != test.payload {
                t.Errorf("Unexpected: TestRespondWithJSON Case #%v: %v", i, output)
				return
			}
		})
	}
}
