package qson

type Stage interface {
	Ensurer
	stageProof()
}

type stage func(m M) M

func (s stage) Ensure(m M) M { return s(initializer().Ensure(m)) }
func (s stage) stageProof()  {}

// AddFields adds new fields to documents. Similar to $project, $addFields
// reshapes each document in the stream; specifically, by adding new fields
// to output documents that contain both the existing fields from the input
// documents and the newly added fields. For more information:
// https://docs.mongodb.com/manual/reference/operator/aggregation/addFields/#pipe._S_addFields
func (aggregation) AddFields(assignments ...Assignment) stage {
	return stage(func(m M) M {
		m["$addFields"] = AGGREGATION.Assignments(assignments...).Ensure(make(M))
		return m
	})
}

// Bucket categorizes incoming documents into groups, called buckets,
// based on a specified expression and bucket boundaries. For more information:
// https://docs.mongodb.com/manual/reference/operator/aggregation/bucket/#pipe._S_bucket
func (aggregation) Bucket(groupBy Expression, boundaries []V, fallback V, output ...Assignment) stage {
	return stage(func(m M) M {
		bucket := M{
			"groupBy":    groupBy.value(),
			"boundaries": boundaries,
		}
		if fallback != "" {
			bucket["default"] = fallback
		}
		if len(output) > 0 {
			bucket["output"] = AGGREGATION.Assignments(output...).Ensure(make(M))
		}
		m["$bucket"] = bucket
		return m
	})
}

// BucketAuto categorizes incoming documents into a specific number of groups,
// called buckets, based on a specified expression. Bucket boundaries are
// automatically determined in an attempt to evenly distribute the documents
// into the specified number of buckets. For more information:
// https://docs.mongodb.com/manual/reference/operator/aggregation/bucketAuto/#pipe._S_bucketAuto
func (aggregation) BucketAuto(groupBy Expression, buckets int, granularity Granularity, output ...Assignment) stage {
	return stage(func(m M) M {
		bucketAuto := M{
			"groupBy": groupBy.value(),
			"buckets": buckets,
		}
		if granularity != "" {
			bucketAuto["granularity"] = granularity
		}
		if len(output) > 0 {
			bucketAuto["output"] = AGGREGATION.Assignments(output...).Ensure(make(M))
		}
		m["$bucketAuto"] = bucketAuto
		return m
	})
}

// CollStats returns statistics regarding a collection or view.
// For more information:
// https://docs.mongodb.com/manual/reference/operator/aggregation/collStats/#pipe._S_collStats
// TODO: think how to define a signature, cause all of four params are optional
func (aggregation) collStats() stage {
	return stage(func(m M) M {
		collStats := M{
			"latencyStats": M{"histograms": true},
			"storageStats": struct{}{},
			"count":        struct{}{},
		}
		m["$collStats"] = collStats
		return m
	})
}

// Count returns a count of the number of documents at this stage of the aggregation pipeline.
// For more information:
// https://docs.mongodb.com/manual/reference/operator/aggregation/count/#pipe._S_count
func (aggregation) Count(output string) stage {
	return stage(func(m M) M {
		m["$count"] = output
		return m
	})
}

// Facet processes multiple aggregation pipelines within a single stage on the same set
// of input documents. Enables the creation of multi-faceted aggregations capable
// of characterizing data across multiple dimensions, or facets, in a single stage.
// For more information:
// https://docs.mongodb.com/manual/reference/operator/aggregation/facet/#pipe._S_facet
func (aggregation) Facet(stages ...AssignedStages) stage {
	return stage(func(m M) M {
		facet := make(M)
		for _, s := range stages {
			s.Ensure(facet)
		}
		m["$facet"] = facet
		return m
	})
}

// Match filters the document stream to allow only matching documents to pass
// unmodified into the next pipeline stage. $match uses standard MongoDB
// queries. For each input document, outputs either one document
// (a match) or zero documents (no match).
func (aggregation) Match(queries ...Query) stage {
	return stage(func(m M) M {
		match := make(M)
		for _, q := range queries {
			q.Ensure(match)
		}
		m["$match"] = match
		return m
	})
}

func (aggregation) Matcher(in ...Query) (match func(...Query) stage) {
	src := make([]Query, 0, len(in))
	src = append(src, in...)
	return func(queries ...Query) stage {
		src = append(src, queries...)
		return AGGREGATION.Match(src...)
	}
}
