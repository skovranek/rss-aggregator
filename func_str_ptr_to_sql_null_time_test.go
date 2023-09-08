package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestStrPtrToSQLNullTime(t *testing.T) {
	now := time.Now().UTC()
	formattedTimeStr := now.Format(time.RFC1123Z)

	removedMonotonicClockTimeStr := now.Format(time.Layout)
	now, err := time.Parse(time.Layout, removedMonotonicClockTimeStr)
	if err != nil {
		t.Errorf("Error: unable to parse time test variable: %v", err)
	}

	var nilPtr *string

	tests := []struct {
		input  *string
		expect sql.NullTime
	}{
		{
			// test zero values
		},
		{
			input: nilPtr,
			expect: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
		},
		{
			input: &formattedTimeStr,
			expect: sql.NullTime{
				Time:  now,
				Valid: true,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output, err := strPtrToSQLNullTime(test.input)
			if err != nil {
				t.Errorf("Error: Test Case #%v: %v", i, err)
			}
			if output != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
