package qson

type LogicalQuery interface {
	Query
	logicalQueryProof()
}

type logicalQuery func(M) M

func (q logicalQuery) Ensure(m M) M       { return q(initializer().Ensure(m)) }
func (q logicalQuery) operatorProof()     {}
func (q logicalQuery) queryProof()        {}
func (q logicalQuery) logicalQueryProof() {}

// And joins query clauses with a logical AND returns all
// documents that match the conditions of both clauses.
func And(queries ...Query) logicalQuery {
	return logicalQuery(func(m M) M {
		and := make([]M, len(queries))
		for i, o := range queries {
			and[i] = o.Ensure(make(M))
		}
		m["$and"] = and
		return m
	})
}

// Not inverts the effect of a logicalQuery expression and returns
// documents that do not match the logicalQuery expression.
func Not(q Query) logicalQuery {
	return logicalQuery(func(m M) M {
		m["$not"] = q.Ensure(make(M))
		return m
	})
}

// Nor joins logicalQuery clauses with a logical NOR returns all
// documents that fail to match both clauses.
func Nor(queries ...Query) logicalQuery {
	return logicalQuery(func(m M) M {
		nor := make([]M, len(queries))
		for i, q := range queries {
			nor[i] = q.Ensure(make(M))
		}
		m["$nor"] = nor
		return m
	})
}

// Or joins logicalQuery clauses with a logical OR returns all
// documents that match the conditions of either clause.
func Or(queries ...Query) logicalQuery {
	return logicalQuery(func(m M) M {
		or := make([]M, len(queries))
		for i, q := range queries {
			or[i] = q.Ensure(make(M))
		}
		m["$or"] = or
		return m
	})
}
