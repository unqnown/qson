package qson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregationPipelineStages(t *testing.T) {
	var ctx = AGGREGATION

	tt := []struct {
		name     string
		m        M
		stage    stage
		expected M
	}{
		{
			name: "$addFields",
			m:    make(M),
			stage: ctx.AddFields(
				ctx.Assign("totalHomework", S("homework")),
				ctx.Assign("totalQuiz", S("quiz")),
			),
			expected: M{
				"$addFields": M{
					"totalHomework": "$homework",
					"totalQuiz":     "$quiz",
				},
			},
		},
		{
			name: "$addFields: nil map",
			m:    nil,
			stage: ctx.AddFields(
				ctx.Assign("totalHomework", S("homework")),
				ctx.Assign("totalQuiz", S("quiz")),
			),
			expected: M{
				"$addFields": M{
					"totalHomework": "$homework",
					"totalQuiz":     "$quiz",
				},
			},
		},
		{
			name: "$bucket",
			m:    make(M),
			stage: ctx.Bucket(
				S("price"),
				[]interface{}{0, 200, 400},
				"Other",
				ctx.Assign("count", ctx.Sum(Value(1))),
				ctx.Assign("title", GROUP.Push(S("title"))),
			),
			expected: M{
				"$bucket": M{
					"groupBy":    "$price",
					"boundaries": []interface{}{0, 200, 400},
					"default":    "Other",
					"output": M{
						"count": M{"$sum": 1},
						"title": M{"$push": "$title"},
					},
				},
			},
		},
		{
			name: "$bucket: without optional fields",
			m:    make(M),
			stage: ctx.Bucket(
				S("price"),
				[]interface{}{0, 200, 400},
				"",
			),
			expected: M{
				"$bucket": M{
					"groupBy":    "$price",
					"boundaries": []interface{}{0, 200, 400},
				},
			},
		},
		{
			name: "$bucketAuto",
			m:    make(M),
			stage: ctx.BucketAuto(
				S("year"),
				5,
				POWERSOF2,
				ctx.Assign("count", ctx.Sum(Value(1))),
				ctx.Assign("year", GROUP.Push(S("year"))),
			),
			expected: M{
				"$bucketAuto": M{
					"groupBy": "$year",
					"buckets": 5,
					"output": M{
						"count": M{"$sum": 1},
						"year":  M{"$push": "$year"},
					},
					"granularity": POWERSOF2,
				},
			},
		},
		{
			name: "$bucketAuto: without optional fields",
			m:    make(M),
			stage: ctx.BucketAuto(
				S("year"),
				5,
				"",
			),
			expected: M{
				"$bucketAuto": M{
					"groupBy": "$year",
					"buckets": 5,
				},
			},
		},
		{
			name:     "$count",
			m:        make(M),
			stage:    ctx.Count("output"),
			expected: M{"$count": "output"},
		},
		{
			name: "$facet",
			m:    make(M),
			stage: ctx.Facet(
				Assign("categorizedByPrice").Stages(
					ctx.Match(
						Exists("price", true),
					),
					ctx.Bucket(
						S("price"),
						[]interface{}{0, 150, 200, 300, 400},
						"Other",
						ctx.Assign("count", ctx.Sum(Value(1))),
						ctx.Assign("titles", GROUP.Push(S("title"))),
					),
				),
			),
			expected: M{
				"$facet": M{
					"categorizedByPrice": []M{
						{
							"$match": M{"price": M{"$exists": true}},
						},
						{
							"$bucket": M{
								"groupBy":    "$price",
								"boundaries": []interface{}{0, 150, 200, 300, 400},
								"default":    "Other",
								"output": M{
									"count":  M{"$sum": 1},
									"titles": M{"$push": "$title"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.stage.Ensure(tc.m)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
