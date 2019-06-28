# WARNING: IN DEVELOPMENT

# qson [![GoDoc](https://godoc.org/github.com/dmytriiandriichuk/qson?status.svg)](https://godoc.org/github.com/dmytriiandriichuk/qson) [![gocover.io](https://gocover.io/_badge/github.com/dmytriiandriichuk/qson)](https://gocover.io/github.com/dmytriiandriichuk/qson)


qson is a simple Go library for wrapping [MongoDB](https://docs.mongodb.com/) procedures 

## Installation

Standard `go get`:

```
$ go get github.com/dmytriiandriichuk/qson
```

## GOALS

- Operators
    - Query and Projection Operators
        - [x] Comparison Query Operators
            - [x] $eq
            - [x] $gt
            - [x] $gte
            - [x] $in
            - [x] $lt
            - [x] $lte
            - [x] $ne
            - [x] $nin
        - [x] Logical Query Operators
            - [x] $and
            - [x] $not
            - [x] $nor
            - [x] $or
        - [x] Element Query Operators
            - [x] $exists
            - [x] $type
        - [ ] Evaluation Query Operators
            - [ ] $expr
            - [ ] $jsonSchema
            - [x] $mod
            - [x] $regex
            - [x] $text
            - [ ] $where
        - [ ] Geospatial Query Operators
            - [ ] $geoIntersects
            - [ ] $geoWithin
            - [ ] $near
            - [ ] $nearSphere
            - [ ] $box
            - [ ] $center
            - [ ] $centerSphere
            - [ ] $geometry
            - [ ] $maxDistance
            - [ ] $minDistance
            - [ ] $polygon
            - [ ] $uniqueDocs
        - [ ] Array Query Operators
            - [ ] $all
            - [ ] $elemMatch
            - [ ] $size
        - [ ] Bitwise Query Operators
            - [ ] $bitsAllClear
            - [ ] $bitsAllSet
            - [ ] $bitsAnyClear
            - [ ] $bitsAnySet
        - [ ] $comment
        - [ ] Projection Operators
            - [x] $ 
            - [ ] $elemMatch 
            - [ ] $meta 
            - [ ] $slice 