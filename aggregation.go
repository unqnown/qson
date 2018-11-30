package qson

// TODO

type Stage interface {
	Ensurer
	stageProof()
}

type stage func(m M) M

func (s stage) Ensure(m M) M { return s(m) }
func (s stage) stageProof()  {}

type match func(m M) M

func (s match) Ensure(m M) M { return s(m) }
func (s match) stageProof()  {}

func Match(queries ...Query) match {
	return match(func(m M) M {
		if m == nil {
			return m
		}
		match := make(M)
		for _, q := range queries {
			q.Ensure(match)
		}
		m["$match"] = match
		return m
	})
}

type matcher struct{ queries []Query }

func Matcher(queries ...Query) *matcher { return &matcher{queries: queries} }

func (s *matcher) Ensure(m M) M { return Match(s.queries...).Ensure(m) }
func (s *matcher) stageProof()  {}
func (s *matcher) Match(queries ...Query) *matcher {
	s.queries = append(s.queries, queries...)
	return s
}

type Aggregation interface {
	Ensurer
	aggregationProof()
}

type aggregation func(M) M

func (a aggregation) Ensure(m M) M      { return a(m) }
func (a aggregation) aggregationProof() {}

func Aggregate(stages ...Stage) aggregation {
	return aggregation(func(m M) M {
		if m == nil {
			return m
		}
		for _, s := range stages {
			s.Ensure(m)
		}
		return m
	})
}
