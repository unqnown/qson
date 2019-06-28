package qson

import "fmt"

type Expression interface {
	Operator
	Assignee
	value() V
	expressionProof()
}

type expression func(M) M

func (e expression) Ensure(m M) M { return e(initializer().Ensure(m)) }
func (e expression) Assign(field string) Assignment {
	return assignment(func(m M) M {
		m[field] = e.Ensure(make(M))
		return m
	})
}
func (e expression) value() V         { return e.Ensure(make(M)) }
func (e expression) operatorProof()   {}
func (e expression) expressionProof() {}

func S(field string) value {
	return value(func() V {
		return fmt.Sprintf("$%s", field)
	})
}

func Value(v V) value {
	return value(func() V {
		return v
	})
}

type value func() V

func (e value) Ensure(m M) M { return m }
func (e value) Assign(field string) Assignment {
	return assignment(func(m M) M {
		m[field] = e.value()
		return m
	})
}
func (e value) value() V         { return e() }
func (e value) operatorProof()   {}
func (e value) expressionProof() {}
