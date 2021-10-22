package routers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/library/db"
	"github.com/library/entities"
	"github.com/library/utils"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	booksCollection = "books"
)

func viewHandler(dal db.DataAccessLayer) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		query := bson.D{{}}
		if val, ok := vars["id"]; ok {
			query = bson.D{{"id", val}}
		}

		books, err := dal.FindBooks(booksCollection, context.Background(), query)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var booksInterface []interface{}
		for _, book := range books {
			booksInterface = append(booksInterface, book)
		}

		// set content type to JSON
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(utils.GenerateResponse(booksInterface, ""))
	})
}

func saveHandler(dal db.DataAccessLayer) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var books []entities.Book
		err = json.Unmarshal(body, &books)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var documents []interface{}
		for _, book := range books {
			document, err := utils.ToDoc(book)
			if err != nil {
				utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
				return
			}
			documents = append(documents, document)
		}

		_, err = dal.InsertBooks(booksCollection, context.Background(), documents)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var booksInterface []interface{}
		for _, book := range books {
			booksInterface = append(booksInterface, book)
		}

		// set content type to JSON
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(utils.GenerateResponse(booksInterface, ""))
	})
}

// BookRouter defines handlers for book specific API endpoints
func BookRouter() *mux.Router {
	client := db.NewMongoDAL()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", viewHandler(client)).Methods("GET")
	router.HandleFunc("/books/{id}", viewHandler(client)).Methods("GET")
	router.HandleFunc("/books", saveHandler(client)).Methods("POST")
	return router
}
