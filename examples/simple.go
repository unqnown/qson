package examples

import (
	"log"

	"github.com/dmytriiandriichuk/qson"
)

func Simple() {
	agg := qson.Aggregate(
		qson.Match(
			qson.Or(
				qson.Same("user_id", "uuid_user_1"),
				qson.Same("user_id", "uuid_user_2"),
			),
			qson.And(
				qson.Or(
					qson.Eq("profession", "software engineer"),
					qson.Eq("profession", "cool guy"),
				),
				qson.Gte("amount", 24),
				qson.Lte("amount", 42),
				qson.Nin("status", []string{"active"}),
			),
			qson.Not(qson.Lte("age", 18)),
		),
	)

	log.Printf("%+v", agg.Ensure(make(qson.M)))
}
