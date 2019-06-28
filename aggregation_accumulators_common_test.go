package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregationAccumulatorsCommon(t *testing.T) {
	var ctx = AGGREGATION

	tt := []struct {
		name        string
		m           M
		accumulator Accumulator
		expected    M
	}{
		{
			name:        "$sum: nil map",
			m:           nil,
			accumulator: ctx.Sum(Value(1)),
			expected:    M{"$sum": 1},
		},
		{
			name:        "$sum: one expression",
			m:           make(M),
			accumulator: ctx.Sum(Value(1)),
			expected:    M{"$sum": 1},
		},
		{
			name:        "$sum: two expressions",
			m:           make(M),
			accumulator: ctx.Sum(Value(1), S("price")),
			expected:    M{"$sum": []interface{}{1, "$price"}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.accumulator.Ensure(tc.m)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
