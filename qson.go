package qson

// M is a convenient alias for a map[string]interface{} map, useful for
// dealing with BSON in a native way.  For instance:
//
//     qson.M{"a": 1, "b": true}
//
// There's no special handling for this type in addition to what's done anyway
// for an equivalent map type.  Elements in the map will be dumped in an
// undefined ordered.
type M = map[string]interface{}

// MS is a convenient alias for a []map[string]interface{} slice of maps.
type MS = []M

// V is concise alias for interface{}, useful for dealing with long
// functions and methods signature.
type V = interface{}

// Ensurer is a base interface which allows to impose mongo syntax requests.
type Ensurer interface {
	// TODO: Ensure(...M) M
	Ensure(M) M
}

type ensurer func(M) M

func (e ensurer) Ensure(m M) M { return e(m) }

func Raw(raw M) ensurer {
	return Merge(
		ensurer(func(m M) M {
			for k, v := range raw {
				m[k] = v
			}
			return m
		}),
	)
}

func intercept(ms ...M) M {
	if len(ms) > 0 {
		return ms[0]
	}
	return make(M)
}

func initializer() ensurer {
	return ensurer(func(m M) M {
		if m == nil {
			return make(M)
		}
		return m
	})
}

// Merge allow to combine several queries based on one field to one expression
func Merge(ensurers ...Ensurer) ensurer {
	return ensurer(func(m M) M {
		var src = make([]M, len(ensurers))
		for i, e := range ensurers {
			sub := make(M)
			e.Ensure(sub)
			src[i] = sub
		}
		for _, sub := range src {
			for field, value := range sub {
				existed, exists := m[field]
				if !exists {
					m[field] = value
					continue
				}
				replacer := make(M)
				switch existed := existed.(type) {
				case M:
					switch value := value.(type) {
					case M:
						Merge(
							ensurer(func(m M) M {
								for k, v := range value {
									m[k] = v
								}
								return m
							}),
							ensurer(func(m M) M {
								for k, v := range existed {
									m[k] = v
								}
								return m
							}),
						).Ensure(replacer)
						m[field] = replacer
					}
				}
			}
		}
		return m
	})
}
