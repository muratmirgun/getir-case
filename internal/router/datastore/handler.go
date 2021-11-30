package datastore

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muratmirgun/getir-case/model"
	"github.com/muratmirgun/getir-case/pkg"
)

func (d *DataStore) InMemory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		d.GetInMemory(w, r)
	case "POST":
		d.SetInMemory(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (d *DataStore) SetInMemory(w http.ResponseWriter, r *http.Request) {
	// Create a new DataHolder
	var input model.DataInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		pkg.Error(err)
		return
	}

	// Set Key-Value pair in memory
	_ = d.Set(input.Key, input.Value)

	// Return Response
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&input)
	if err != nil {
		pkg.Error(err)
		return
	}
}

func (d *DataStore) GetInMemory(w http.ResponseWriter, r *http.Request) {
	// Read Key from Header
	HeaderKey := r.Header.Get("key")

	// Get Key-Value Pair from DataHolder
	value, err := d.Get(HeaderKey)
	if err != nil {
		_, err := fmt.Fprintf(w, "%+v", err.Error())
		if err != nil {
			pkg.Error(err)
			return
		}
	} else {
		out := model.DataInput{Key: HeaderKey, Value: value}
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			pkg.Error(err)
			return
		}
	}
}
