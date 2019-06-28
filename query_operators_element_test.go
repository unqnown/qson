package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryOperators_Element(t *testing.T) {
	tt := []struct {
		name     string
		query    Query
		expected M
	}{
		{
			name:  "exists",
			query: Exists("user_id", true),
			expected: M{
				"user_id": M{"$exists": true},
			},
		},
		{
			name:  "type",
			query: Type("user_id", 0x02),
			expected: M{
				"user_id": M{"$type": []byte{0x02}},
			},
		},
		{
			name:  "regex",
			query: Regex("user_id", "[A-Z]", "i", "m"),
			expected: M{
				"user_id": M{"$regex": "[A-Z]", "$options": "im"},
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
