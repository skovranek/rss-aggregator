package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestStrPtrToSQLNullStr(t *testing.T) {
	str := "this is a string"
	var nilPtr *string

	tests := []struct {
		input  *string
		expect sql.NullString
	}{
		{
			input: nilPtr,
			expect: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			input: &str,
			expect: sql.NullString{
				String: str,
				Valid:  true,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestStrPtrToSQLNullStr Case #%v:", i), func(t *testing.T) {
			output, _ := strPtrToSQLNullStr(test.input)
			if output != test.expect {
                t.Errorf("Unexpected: TestStrPtrToSQLNullStr:\n%v", output)
				return
			}
		})
	}
}
