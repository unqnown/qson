package qson

// Sum returns a sum of numerical values. Ignores non-numeric values.
func (aggregation) Sum(expressions ...Expression) accumulator {
	return accumulator(func(m M) M {
		if len(expressions) == 1 {
			m["$sum"] = expressions[0].value()
			return m
		}
		sum := make([]V, len(expressions))
		for i, exp := range expressions {
			sum[i] = exp.value()
		}
		m["$sum"] = sum
		return m
	})
}
