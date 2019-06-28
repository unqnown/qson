package qson

type group string

const GROUP group = "group"

// Sum returns a sum of numerical values. Ignores non-numeric values.
func (group) Sum(expressions ...Expression) accumulator { return AGGREGATION.Sum(expressions...) }

// Push returns an array of expression values for each group.
func (group) Push(expr Expression) accumulator {
	return accumulator(func(m M) M {
		m["$push"] = expr.value()
		return m
	})
}
