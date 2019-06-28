package qson

// Abs returns the absolute value of a number.
func (aggregation) Abs(number Expression) expression {
	return func(m M) M {
		m["$abs"] = number.value()
		return m
	}
}
