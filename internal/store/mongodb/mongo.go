package mongodb

import (
	"context"
	"fmt"
	"getir-case/internal/store"
	"getir-case/model"
	"getir-case/pkg"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	collection *mongo.Collection
}

func MongoInstance() store.DataManager {
	Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(Ctx, options.Client().ApplyURI("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"))
	if err != nil {
		pkg.Panic(err)
	}
	session, err := client.StartSession()
	if err != nil {
		pkg.Panic(err)
	}
	database := session.Client().Database("getir-case-study")

	recordsCollection := database.Collection("records")
	return &mongodb{collection: recordsCollection}
}

var mongoIns = MongoInstance()

func MongoMgr() store.DataManager { return mongoIns }

func (m *mongodb) Retrieve(input interface{}) (out interface{}, err error) {
	// Create Variables
	var rData []bson.M
	var Resp model.MongoResponse
	var Req model.MongoRequest

	Req, _ = input.(model.MongoRequest)

	// Default Data's
	Resp.Code = http.StatusBadRequest
	Resp.Records = rData

	sd, err := time.Parse("2006-01-02", Req.StartDate)
	if err != nil {
		Resp.Msg = err.Error()
		pkg.Error(err)
		return Resp, err
	}

	ed, err := time.Parse("2006-01-02", Req.EndDate)
	if err != nil {
		Resp.Msg = err.Error()
		pkg.Error(err)
		return Resp, err
	}

	// References for pipeline stages and some Helps from here
	// https://docs.mongodb.com/manual/reference/operator/aggregation-pipeline/
	// https://docs.mongodb.com/manual/core/aggregation-pipeline/
	//  https://docs.mongodb.com/manual/reference/operator/aggregation/match/
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": sd,
					"$lt": ed,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": Req.MinCount,
					"$lt": Req.MaxCount,
				},
			},
		},
	}

	cursor, err := m.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		Resp.Msg = err.Error()
		pkg.Error(err)
		return Resp, err
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &rData); err != nil {
		Resp.Msg = err.Error()
		pkg.Error(err)
		return Resp, err
	}

	if len(rData) > 0 {
		Resp.Code = 0
		Resp.Msg = "Success"
		Resp.Records = rData
		return Resp, nil
	}

	Resp.Code = http.StatusNoContent
	Resp.Msg = "Data not Found"
	Resp.Records = rData
	err = fmt.Errorf("Data Not found")
	return Resp, err
}
