package examples

import (
	"encoding/json"
	"fmt"

	"github.com/dmytriiandriichuk/qson"
)

func Simple()  {
	agg := qson.Aggregate(
		qson.Match(
			qson.Or(
				qson.Eq("profession", "software engineer"),
				qson.Eq("profession", "cool guy"),
			),
			qson.And(
				qson.Gte("experience", 24),
				qson.Lte("experience", 42),
				qson.Nin("status", []string{"active"}),
			),
			qson.Not(qson.Lte("age", 18)),
		),
	)

	j, _ := json.MarshalIndent(agg.Ensure(make(qson.M)), "", "	")
	fmt.Printf("%s", string(j))
}
