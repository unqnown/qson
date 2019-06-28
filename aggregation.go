package qson

type aggregation string

// AGGREGATION is exported namespace which provides aggregation expressions.
const AGGREGATION aggregation = "aggregation"

// TODO: implement aggregator

// There is no way to implement ensurer interface. Sorry
// Aggregate allows to build aggregation pipeline.
func (aggregation) Aggregate(stages ...Stage) (pipeline MS) {
	pipeline = make(MS, len(stages))
	for i, s := range stages {
		pipeline[i] = s.Ensure(make(M))
	}
	return pipeline
}

func Namespaces() (agg aggregation, prj project, grp group) { return AGGREGATION, PROJECT, GROUP }

// A string that specifies the preferred number series to use to ensure
// that the calculated boundary edges end on preferred round numbers or
// their powers of 10.
type Granularity = string

// Possible values of granularity
const (
	R5        Granularity = "R5"
	R10       Granularity = "R10"
	R20       Granularity = "R20"
	R40       Granularity = "R40"
	R80       Granularity = "R80"
	_1_2_5    Granularity = "1-2-5"
	E6        Granularity = "E6"
	E12       Granularity = "E12"
	E24       Granularity = "E24"
	E48       Granularity = "E48"
	E96       Granularity = "E96"
	E192      Granularity = "E192"
	POWERSOF2 Granularity = "POWERSOF2"
)
