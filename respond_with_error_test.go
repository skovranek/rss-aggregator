package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http/httptest"
    "testing"
)

func TestRespondWithError(t *testing.T) {
	tests := []struct {
        errMsg string
        statusCode int
	}{
        {}, // zero values
        {
            errMsg: "this is a string",
            statusCode: 200,
        },
	}

	for i, test := range tests {
        t.Run(fmt.Sprintf("TestRespondWithError Case #%v:", i), func(t *testing.T) {
            w := httptest.NewRecorder()

            respondWithError(w, test.statusCode, test.errMsg)

            res := w.Result()
            body, err := io.ReadAll(res.Body)
            if err != nil {
                t.Errorf("Error: TestRespondWithError: io.ReadAll(res.Body): %v", err)
                return
            }

            if res.StatusCode != test.statusCode {
                if test.statusCode > 99 && test.statusCode < 600 {
                    t.Errorf("Unexpected: TestRespondWithError Case #%v: %v", i, res.StatusCode)
                    return
                }
			}

			if res.Header.Get("Content-Type") != "application/json" {
                t.Errorf("Unexpected: TestRespondWithError Case #%v: %v", i, res.Header.Get("Content-Type"))
				return
            }

            output := struct{
                Error any `json:"error"`
            }{}

			err = json.Unmarshal(body, &output)
			if err != nil {
                t.Errorf("Error: TestRespondWithError Case #%v: %v", i, err)
				return
			}

            if output.Error != test.errMsg {
                t.Errorf("Unexpected: TestRespondWithError Case #%v: %v", i, output)
				return
			}
		})
	}
}
