package qson

type Assignment interface {
	Ensurer
	assignmentProof()
}

type Assignee interface {
	Assign(string) Assignment
}

type assignment func(m M) M

func (a assignment) Ensure(m M) M     { return a(initializer().Ensure(m)) }
func (a assignment) assignmentProof() {}

func (aggregation) Assign(field string, assignee Assignee) assignment {
	return assignment(func(m M) M {
		return assignee.Assign(field).Ensure(m)
	})
}

func (aggregation) Assignments(assignments ...Assignment) assignment {
	return assignment(func(m M) M {
		for _, a := range assignments {
			m = a.Ensure(m)
		}
		return m
	})
}

type assigner func() string

func (a assigner) Stages(stages ...Stage) assignedStages {
	return assignedStages(func(m M) M {
		m[a()] = AGGREGATION.Aggregate(stages...)
		return m
	})
}

// TODO
func (a assigner) Expression(expr Expression)    {}
func (a assigner) Accumulator(accum Accumulator) {}

func Assign(field string) assigner {
	return assigner(func() string {
		return field
	})
}

type AssignedStages interface {
	Assignment
	assignedStagesProof()
}

type assignedStages func(m M) M

func (a assignedStages) Ensure(m M) M         { return a(initializer().Ensure(m)) }
func (a assignedStages) assignmentProof()     {}
func (a assignedStages) assignedStagesProof() {}
