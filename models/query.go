package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"strconv"
	"strings"
)

type Query struct {
	Keyword			string			`json:"keyword"`
	Sort 			string			`json:"sort"`
	Limit 			string			`json:"limit"`
	Page 			string			`json:"page"`
}

type Pagination struct {
	Count 			int64 			`json:"count"`
	Page 			int64			`json:"page"`
	Pages 			int64			`json:"pages"`
	Skip 			int64			`json:"skip"`
	Limit 			int64 			`json:"limit"`
}

type Result struct {
	Data			interface{}		`json:"data"`
	Pagination		Pagination		`json:"pagination"`
	Query 			Query			`json:"query"`
}

func (q Query) GetPagination(count int64) Pagination {
	pages := int64(0)
	page := int64(1)
	limit := int64(0)

	if q.Limit != "" {
		limit, _ = strconv.ParseInt(q.Limit, 10, 64)
		pages = int64(math.Round(float64(count) / float64(limit)))
	}

	if q.Page != "" {
		page, _ = strconv.ParseInt(q.Page, 10, 64)
	}

	skip := page * limit - limit

	return Pagination{
		Count: count,
		Page: page,
		Pages: pages,
		Skip: skip,
		Limit: limit,
	}
}

func (q Query) GetOptions() *options.FindOptions {
	opts := options.Find()

	limit, _ := strconv.ParseInt(q.Limit, 10, 64)
	opts.SetLimit(limit)

	if q.Sort != "" {
		sorts := bson.D{}
		splittedFields := strings.Split(q.Sort, "|")
		for _, val := range splittedFields {
			splitted := strings.Split(val, ",")

			key := splitted[0]
			value, _ := strconv.ParseInt(splitted[1], 10, 64)

			sorts = append(sorts, bson.E{Key: key, Value: value})
		}

		opts.SetSort(sorts)
	}

	page := int64(1)
	if q.Page != "" {
		page, _ = strconv.ParseInt(q.Page, 10, 64)
	}

	skip := page * limit - limit
	opts.SetSkip(skip)

	return opts
}

func (q Query) GetQueryFind() bson.M {
	query := bson.M{}
	orFilter := bson.A{}

	if q.Keyword != "" {
		regex := primitive.Regex{Pattern: "^" + q.Keyword, Options: "i"}
		orFilter = append(orFilter, bson.M{ "name": regex})
	}

	if len(orFilter) > 0 {
		query["$or"] = orFilter
	}

	return query
}
