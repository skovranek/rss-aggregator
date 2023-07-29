package database

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
)

func (q *Queries) TestGetAllFeeds(t *testing.T) {
	ctx := context.Background()

	t.Run("TestGetAllFeeds:", func(t *testing.T) {
		output, err := q.GetAllFeeds(ctx)
		if err != nil {
			t.Errorf("Error: %v\n", err)
			return
		}

		if output == nil {
			t.Errorf("Unexpected:\n%v", output)
			return
		}
	})
}
