package qson

type ComparisonQuery interface {
	Query
	comparisonQueryProof()
}

type comparisonQuery func(M) M

func (q comparisonQuery) Ensure(m M) M          { return q(initializer().Ensure(m)) }
func (q comparisonQuery) operatorProof()        {}
func (q comparisonQuery) queryProof()           {}
func (q comparisonQuery) comparisonQueryProof() {}

// Same matches values that are equal to a specified value.
func Same(field, value string) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = value
		return m
	})
}

// Eq matches values that are equal to a specified value.
func Eq(field string, value interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$eq": value}
		return m
	})
}

// Gt matches values that are greater than a specified value.
func Gt(field string, value interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$gt": value}
		return m
	})
}

// Gte matches values that are greater than or equal to a specified value.
func Gte(field string, value interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$gte": value}
		return m
	})
}

// In matches any of the values specified in an array.
func In(field string, values interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$in": values}
		return m
	})
}

// Lt matches values that are less than a specified value.
func Lt(field string, value interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$lt": value}
		return m
	})
}

// Lte matches values that are less than or equal to a specified value.
func Lte(field string, value interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$lte": value}
		return m
	})
}

// Ne matches all values that are not equal to a specified value.
func Ne(field string, value interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$ne": value}
		return m
	})
}

// Nin matches none of the values specified in an array.
func Nin(field string, values interface{}) comparisonQuery {
	return comparisonQuery(func(m M) M {
		m[field] = M{"$nin": values}
		return m
	})
}
