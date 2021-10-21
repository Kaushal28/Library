package utils

import (
	"testing"

	"github.com/library/entities"
)

func TestGenerateResponse(t *testing.T) {
	response := GenerateResponse(nil, "sample_error")

	if response.Data == nil {
		t.Errorf("Got nil valued data data.")
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