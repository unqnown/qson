# qson [![GoDoc](https://godoc.org/github.com/dmytriiandriichuk/qson?status.svg)](https://godoc.org/github.com/dmytriiandriichuk/qson)

qson is a simple Go library for wrapping [MongoDB](https://docs.mongodb.com/) procedures 

## Installation

Standard `go get`:

```
$ go get github.com/dmytriiandriichuk/qson
```

## Usage & Example

A quick code example is shown below:

```go
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
// Output:
// {
//    "$match": {
//        "$and": [
//            {
//                "experience": {
//                    "$gte": 24
//                }
//            },
//            {
//                "experience": {
//                    "$lte": 42
//                }
//            },
//            {
//                "status": {
//                    "$nin": [
//                        "active"
//                    ]
//                }
//            }
//        ],
//        "$not": {
//            "age": {
//                "$lte": 18
//            }
//        },
//        "$or": [
//            {
//                "profession": {
//                    "$eq": "software engineer"
//                }
//            },
//            {
//                "profession": {
//                    "$eq": "cool guy"
//                }
//            }
//        ]
//    }
//}
```