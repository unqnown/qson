package qson

type Query interface {
	Operator
	queryProof()
}

// query is a general realization of Query operator.
type query func(M) M

func (q query) Ensure(m M) M   { return q(m) }
func (q query) operatorProof() {}
func (q query) queryProof()    {}

// Queries performs group overlaying queries on request form.
// Queries for the same field will conflict and each new query
// for the same field will be override previous one.
func Queries(queries ...Query) query {
	return query(func(m M) M {
		for _, q := range queries {
			q.Ensure(m)
		}
		return m
	})
}
