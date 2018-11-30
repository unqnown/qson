package qson

type LogicalQuery interface {
	Query
	logicalQueryProof()
}

type logicalQuery func(M) M

func (q logicalQuery) Ensure(m M) M       { return q(m) }
func (q logicalQuery) operatorProof()     {}
func (q logicalQuery) queryProof()        {}
func (q logicalQuery) logicalQueryProof() {}

// And joins query clauses with a logical AND returns all
// documents that match the conditions of both clauses.
func And(queries ...Query) logicalQuery {
	return logicalQuery(func(m M) M {
		if m == nil {
			return m
		}
		and := make([]M, len(queries))
		for i, o := range queries {
			sub := make(M)
			o.Ensure(sub)
			and[i] = sub
		}
		m["$and"] = and
		return m
	})
}

// Not inverts the effect of a logicalQuery expression and returns
// documents that do not match the logicalQuery expression.
func Not(q Query) logicalQuery {
	return logicalQuery(func(m M) M {
		if m == nil {
			return m
		}
		exp := make(M)
		q.Ensure(exp)
		m["$not"] = exp
		return m
	})
}

// Nor joins logicalQuery clauses with a logical NOR returns all
// documents that fail to match both clauses.
func Nor(queries ...Query) logicalQuery {
	return logicalQuery(func(m M) M {
		if m == nil {
			return m
		}
		nor := make([]M, len(queries))
		for i, q := range queries {
			sub := make(M)
			q.Ensure(sub)
			nor[i] = sub
		}
		m["$nor"] = nor
		return m
	})
}

// Or joins logicalQuery clauses with a logical OR returns all
// documents that match the conditions of either clause.
func Or(queries ...Query) logicalQuery {
	return logicalQuery(func(m M) M {
		if m == nil {
			return m
		}
		and := make([]M, len(queries))
		for i, q := range queries {
			sub := make(M)
			q.Ensure(sub)
			and[i] = sub
		}
		m["$or"] = and
		return m
	})
}
