package qson

type ElementQuery interface {
	Query
	elementQueryProof()
}

type elementQuery func(M) M

func (q elementQuery) Ensure(m M) M       { return q(m) }
func (q elementQuery) operatorProof()     {}
func (q elementQuery) queryProof()        {}
func (q elementQuery) elementQueryProof() {}

// Exists matches documents that have the specified field.
func Exists(field string, exists bool) elementQuery {
	return elementQuery(func(m M) M {
		if m == nil {
			return m
		}
		m[field] = M{"$exists": exists}
		return m
	})
}

// Type selects documents if a field is of the specified type.
func Type(field string, types ...byte) elementQuery {
	return elementQuery(func(m M) M {
		if m == nil {
			return m
		}
		m[field] = M{"$type": types}
		return m
	})
}
