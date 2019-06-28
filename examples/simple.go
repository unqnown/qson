package examples

import (
	"encoding/json"
	"fmt"

	"github.com/unqnown/qson"
)

func Simple() {
	agg := qson.AGGREGATION
	aggregation := agg.Aggregate(
		agg.Match(
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

	j, _ := json.MarshalIndent(aggregation, "", "	")
	fmt.Printf("%s", string(j))
}
