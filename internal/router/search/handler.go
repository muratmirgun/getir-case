package search

import (
	"encoding/json"
	"getir-case/internal/store/mongodb"
	"getir-case/model"
	"getir-case/pkg"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoDB struct{}

func (m *MongoDB) ServeMongo(rw http.ResponseWriter, request *http.Request) {
	var result interface{}
	var mongoResponse model.MongoResponse
	var data []bson.M

	mongoResponse.Code = http.StatusBadRequest
	mongoResponse.Records = data

	if request.Method != "POST" {
		mongoResponse.Msg = "Method not allowed"
		rw.WriteHeader(500)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			pkg.Error(err)
			return
		}
		return
	}
	if nil == request.Body {
		mongoResponse.Msg = "No request content to process"
		rw.WriteHeader(500)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			pkg.Error(err)
			return
		}
		return
	}

	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		mongoResponse.Msg = err.Error()
		pkg.Error(err)
		rw.WriteHeader(500)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			pkg.Error(err)
			return
		}
		return
	}

	var content model.MongoRequest

	if err = json.Unmarshal(body, &content); err != nil {
		rw.WriteHeader(500)
		mongoResponse.Msg = err.Error()
		pkg.Error(err)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			pkg.Error(err)
			return
		}
		return
	}

	result, err = mongodb.MongoMgr().Retrieve(content)
	if err != nil {
		pkg.Error(err)
		return
	}
	err = json.NewEncoder(rw).Encode(result)
	if err != nil {
		pkg.Error(err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
