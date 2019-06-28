package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpression(t *testing.T) {
	tt := []struct {
		name       string
		expression Expression
		expected   M
	}{
		{
			name:       "abs field",
			expression: AGGREGATION.Abs(S("value")),
			expected: M{
				"$abs": "$value",
			},
		},
		{
			name:       "abs number",
			expression: AGGREGATION.Abs(Value(42)),
			expected: M{
				"$abs": 42,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.expression.Ensure(make(M)))
		})
	}
}
