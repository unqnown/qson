package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryOperators_Logical(t *testing.T) {
	tt := []struct {
		name     string
		query    Query
		expected M
	}{
		{
			name: "and",
			query: And(
				Same("user_id", "uuid_user"),
				In("status", []string{"active", "pending"}),
			),
			expected: M{
				"$and": []M{
					{"user_id": "uuid_user"},
					{"status": M{"$in": []string{"active", "pending"}}},
				},
			},
		},
		{
			name: "not",
			query: Not(
				Same("user_id", "uuid_user"),
			),
			expected: M{
				"$not": M{"user_id": "uuid_user"},
			},
		},
		{
			name: "nor",
			query: Nor(
				Same("user_id", "uuid_user"),
				In("status", []string{"active", "pending"}),
			),
			expected: M{
				"$nor": []M{
					{"user_id": "uuid_user"},
					{"status": M{"$in": []string{"active", "pending"}}},
				},
			},
		},
		{
			name: "or",
			query: Or(
				Same("user_id", "uuid_user"),
				In("status", []string{"active", "pending"}),
			),
			expected: M{
				"$or": []M{
					{"user_id": "uuid_user"},
					{"status": M{"$in": []string{"active", "pending"}}},
				},
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
