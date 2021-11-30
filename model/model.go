package model

import "go.mongodb.org/mongo-driver/bson"

type MongoRequest struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	MinCount  float64 `json:"minCount"`
	MaxCount  float64 `json:"maxCount"`
}

type MongoResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []bson.M `json:"records"`
}

type DataInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
