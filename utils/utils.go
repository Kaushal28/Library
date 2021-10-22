package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Kaushal28/library/entities"
	"go.mongodb.org/mongo-driver/bson"
)

// ToDoc converts object of "any type" to bson document
func ToDoc(v interface{}) (*bson.D, error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	var doc *bson.D
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// JSONError returns errors in JSON format. It's analogous to http.Error()
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

// GenerateResponse generates standard response from given error string and list of data objects
func GenerateResponse(data []interface{}, err string) entities.Response {
	if data == nil {
		data = make([]interface{}, 0)
	}

	var error entities.Error
	if err != "" {
		error = entities.Error{Message: err}
	}
	return entities.Response{Error: error, Data: data}
}
