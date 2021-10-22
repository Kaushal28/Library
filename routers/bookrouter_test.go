package routers

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Kaushal28/library/entities"
)

type mockDAL struct{}

// implementing data access layer's methods on MongoDAL to satisfy DataAccessLayer interface

// Findbooks queries mongodb with given query to fetch relevant books
func (m *mockDAL) FindBooks(coll string, ctx context.Context, query interface{}) ([]entities.Book, error) {
	// return dummy data
	return []entities.Book{{Title: "one"}, {Title: "two"}}, nil
}

// InsertBooks inserts given list of books into mongodb
func (m *mockDAL) InsertBooks(coll string, ctx context.Context, documents []interface{}) (interface{}, error) {
	return nil, nil
}

func TestViewHandler(t *testing.T) {
	// get the handler by injecting dependency
	handler := viewHandler(&mockDAL{})

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)

	handler(rw, req)

	res := rw.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected error to be nil got %v", err)
	}

	expected := "{\"error\":{\"message\":\"\"},\"data\":[{\"id\":\"\",\"title\":\"one\",\"author\":\"\",\"cost\":0,\"year\":0,\"publisher\":\"\"},{\"id\":\"\",\"title\":\"two\",\"author\":\"\",\"cost\":0,\"year\":0,\"publisher\":\"\"}]}"
	if strings.TrimSpace(string(data)) != expected {
		t.Errorf("Expected data: %v, found %v", "", strings.TrimSpace(string(data)))
	}
}

func TestInsertHandler(t *testing.T) {
	// get the handler by injecting dependency
	handler := viewHandler(&mockDAL{})

	rw := httptest.NewRecorder()
	body := "{\"error\":{\"message\":\"\"},\"data\":[{\"id\":\"\",\"title\":\"one\",\"author\":\"\",\"cost\":0,\"year\":0,\"publisher\":\"\"},{\"id\":\"\",\"title\":\"two\",\"author\":\"\",\"cost\":0,\"year\":0,\"publisher\":\"\"}]}"
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(body))

	handler(rw, req)
	res := rw.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected error to be nil got %v", err)
	}

	if strings.TrimSpace(string(data)) != body {
		t.Errorf("Expected data: %v, found %v", "", strings.TrimSpace(string(data)))
	}
}
