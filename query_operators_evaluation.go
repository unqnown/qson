package qson

import "fmt"

type EvaluationQuery interface {
	Query
	evaluationQueryProof()
}

type evaluationQuery func(M) M

func (q evaluationQuery) Ensure(m M) M          { return q(initializer().Ensure(m)) }
func (q evaluationQuery) operatorProof()        {}
func (q evaluationQuery) queryProof()           {}
func (q evaluationQuery) evaluationQueryProof() {}

// Regex selects documents where values match a specified
// regular expression.
func Regex(field, pattern string, options ...string) evaluationQuery {
	return evaluationQuery(func(m M) M {
		m[field] = M{
			"$regex": pattern,
			"$options": fmt.Sprint(
				func() (opts []interface{}) {
					opts = make([]interface{}, len(options))
					for i, o := range options {
						opts[i] = o
					}
					return
				}()...,
			),
		}
		return m
	})
}

// Text performs text search.
func Text(text string) evaluationQuery {
	return evaluationQuery(func(m M) M {
		m["$text"] = M{"$search": text}
		return m
	})
}

// Mod performs a modulo operation on the value of a field and
// selects documents with a specified result.
func Mod(field string, divisor, remainder int64) evaluationQuery {
	return evaluationQuery(func(m M) M {
		m[field] = M{"$mod": []int64{divisor, remainder}}
		return m
	})
}
