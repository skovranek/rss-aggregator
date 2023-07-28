package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestStrPtrToSQLNullStr(t *testing.T) {
	testStr := "this is a string"
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
			input: &testStr,
			expect: sql.NullString{
				String: testStr,
				Valid:  true,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output, _ := strPtrToSQLNullStr(test.input)
			if output != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
