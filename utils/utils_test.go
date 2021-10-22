package utils

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/library/entities"
)

func TestGenerateResponse(t *testing.T) {
	response := GenerateResponse(nil, "sample_error")

	if response.Data == nil {
		t.Errorf("Got nil valued data.")
	}

	if response.Error.Message != "sample_error" {
		t.Error("Got incorrect error message.")
	}

	book := entities.Book{}
	var bookInterface []interface{}
	bookInterface = append(bookInterface, book)

	response = GenerateResponse(bookInterface, "")
	book, ok := response.Data.([]interface{})[0].(entities.Book)

	if !ok {
		t.Error("Got invalid Book data.")
	}
}

func TestJSONError(t *testing.T) {
    rw := httptest.NewRecorder()
	JSONError(rw, entities.Error{Message: "This is a test error!"}, 500)
	res := rw.Result()
    defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }
    if strings.TrimSpace(string(data)) != "{\"message\":\"This is a test error!\"}" {
        t.Errorf("expected error in JSON format, got %v", string(data))
    }
}