package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_Queries(t *testing.T) {
	tt := []struct {
		name     string
		queries  []Query
		expected M
	}{
		{
			name: "default",
			queries: []Query{
				In("status", []string{"active", "pending"}),
				Same("user_id", "uuid_user"),
			},
			expected: M{
				"user_id": "uuid_user",
				"status":  M{"$in": []string{"active", "pending"}},
			},
		},
	}

	for _, tc := range tt {
		var actual = make(M)
		Queries(tc.queries...).Ensure(actual)
		assert.Equal(t, tc.expected, actual)
	}
}
