package qson

type Accumulator interface {
	Expression
	accumulatorProof()
}

type accumulator func(m M) M

func (a accumulator) Ensure(m M) M { return a(initializer().Ensure(m)) }
func (a accumulator) Assign(field string) Assignment {
	return assignment(func(m M) M {
		m[field] = a.Ensure(make(M))
		return m
	})
}
func (a accumulator) value() V          { return a.Ensure(make(M)) }
func (a accumulator) operatorProof()    {}
func (a accumulator) expressionProof()  {}
func (a accumulator) accumulatorProof() {}
