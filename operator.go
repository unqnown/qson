package qson

type Operator interface {
	Ensurer
	operatorProof()
}

// operator is a general realization of Operator.
type operator func(M) M

func (o operator) Ensure(m M) M   { return o(m) }
func (o operator) operatorProof() {}

func Operators(operators ...Operator) operator {
	return operator(func(m M) M {
		for _, o := range operators {
			o.Ensure(m)
		}
		return m
	})
}
