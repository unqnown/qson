package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionOperatorsArithmetic(t *testing.T) {
	var ctx = AGGREGATION

	tt := []struct {
		name       string
		m          M
		expression Expression
		expected   M
	}{
		{
			name:       "$abs: bind",
			m:          make(M),
			expression: ctx.Abs(S("price")),
			expected:   M{"$abs": "$price"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.expression.Ensure(tc.m)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
