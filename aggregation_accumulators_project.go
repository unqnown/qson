package qson

type project string

const PROJECT project = "project"

// Sum returns a sum of numerical values. Ignores non-numeric values.
func (project) Sum(expressions ...Expression) accumulator { return AGGREGATION.Sum(expressions...) }
