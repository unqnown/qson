package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryOperators_Comparison(t *testing.T) {
	tt := []struct {
		name     string
		query    Query
		expected M
	}{
		{
			name:     "same",
			query:    Same("user_id", "uuid_user"),
			expected: M{"user_id": "uuid_user"},
		},
		{
			name:  "eq",
			query: Eq("amount", 42),
			expected: M{
				"amount": M{"$eq": 42},
			},
		},
		{
			name:  "gt",
			query: Gt("amount", 42),
			expected: M{
				"amount": M{"$gt": 42},
			},
		},
		{
			name:  "gte",
			query: Gte("amount", 42),
			expected: M{
				"amount": M{"$gte": 42},
			},
		},
		{
			name:  "in",
			query: In("status", []string{"active", "pending"}),
			expected: M{
				"status": M{"$in": []string{"active", "pending"}},
			},
		},
		{
			name:  "lt",
			query: Lt("amount", 42),
			expected: M{
				"amount": M{"$lt": 42},
			},
		},
		{
			name:  "lte",
			query: Lte("amount", 42),
			expected: M{
				"amount": M{"$lte": 42},
			},
		},
		{
			name:  "ne",
			query: Ne("amount", 42),
			expected: M{
				"amount": M{"$ne": 42},
			},
		},
		{
			name:  "nin",
			query: Nin("status", []string{"active", "pending"}),
			expected: M{
				"status": M{"$nin": []string{"active", "pending"}},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var actual = make(M)
			tc.query.Ensure(actual)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
