package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsurer_Merge(t *testing.T) {
	tt := []struct {
		name     string
		ensurers []Ensurer
		expected M
	}{
		{
			name: "easy",
			ensurers: []Ensurer{
				Gt("amount", 24),
				Lt("amount", 42),
			},
			expected: M{
				"amount": M{"$gt": 24, "$lt": 42},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var actual = make(M)
			Merge(tc.ensurers...).Ensure(actual)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
