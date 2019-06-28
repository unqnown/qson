package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryOperators_Evaluation(t *testing.T) {
	tt := []struct {
		name     string
		query    Query
		expected M
	}{
		{
			name:  "$regex",
			query: Regex("user_id", "pattern", "options"),
			expected: M{
				"user_id": M{"$regex": "pattern", "$options": "options"},
			},
		},
		{
			name:  "$text",
			query: Text("some text"),
			expected: M{
				"$text": M{"$search": "some text"},
			},
		},
		{
			name:  "$mod",
			query: Mod("amount", 3, 0),
			expected: M{
				"amount": M{"$mod": []int64{3, 0}},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.query.Ensure(make(M))
			assert.Equal(t, tc.expected, actual)
		})
	}
}
