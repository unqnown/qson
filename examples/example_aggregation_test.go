package examples

import (
	"encoding/json"
	"fmt"

	"github.com/unqnown/qson"
)

func ExampleAggregation() {
	agg, _, _ := qson.Namespaces()

	match := agg.Matcher()

	match(
		qson.Lte("lol", 2),
	)

	aggregation := agg.Aggregate(
		match(
			qson.Eq("smth", "lol"),
		),
	)

	j, _ := json.MarshalIndent(aggregation, "", "	")
	fmt.Printf("%s", string(j))
	// Output: [
	//	{
	//		"$match": {
	//			"lol": {
	//				"$lte": 2
	//			},
	//			"smth": {
	//				"$eq": "lol"
	//			}
	//		}
	//	}
	//]
}
