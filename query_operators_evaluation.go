package qson

type EvaluationQuery interface {
	Query
	evaluationQueryProof()
}

type evaluationQuery func(M) M

func (q evaluationQuery) Ensure(m M) M          { return q(m) }
func (q evaluationQuery) operatorProof()        {}
func (q evaluationQuery) queryProof()           {}
func (q evaluationQuery) evaluationQueryProof() {}

// Regex selects documents where values match a specified
// regular expression.
func Regex(field, pattern, options string) evaluationQuery {
	return evaluationQuery(func(m M) M {
		if m == nil {
			return m
		}
		m[field] = M{"$regex": pattern, "$options": options}
		return m
	})
}

// Text performs text search.
func Text(text string) evaluationQuery {
	return evaluationQuery(func(m M) M {
		if m == nil {
			return m
		}
		m["$text"] = M{"$search": text}
		return m
	})
}

// Mod performs a modulo operation on the value of a field and
// selects documents with a specified result.
func Mod(field string, divisor, remainder int64) evaluationQuery {
	return evaluationQuery(func(m M) M {
		if m == nil {
			return m
		}
		m[field] = M{"$mod": []int64{divisor, remainder}}
		return m
	})
}
