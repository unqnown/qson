package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsurer(t *testing.T) {
	tt := []struct {
		name     string
		ensurer  Ensurer
		expected M
	}{
		{
			name: "merge",
			ensurer: Merge(
				Gt("amount", 24),
				Lt("amount", 42),
			),
			expected: M{
				"amount": M{"$gt": 24, "$lt": 42},
			},
		},
		{
			name: "raw",
			ensurer: Raw(
				M{
					"field_1": "value_1",
					"field_2": 2,
				},
			),
			expected: M{
				"field_1": "value_1",
				"field_2": 2,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.ensurer.Ensure(make(M)))
		})
	}
}
