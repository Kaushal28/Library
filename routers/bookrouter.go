package routers

import (
	"context"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/library/db"
	"github.com/library/entities"
	"github.com/library/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	booksCollection = "books"
)

func viewHandler(client *mongo.Client) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		coll := client.Database(db.Database).Collection(booksCollection)

		vars := mux.Vars(r)
		query := bson.D{{}}
		if val, ok := vars["id"]; ok {
			query = bson.D{{"id", val}}
		}

		cursor, err := coll.Find(context.TODO(), query)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var books = make([]entities.Book, 0)
		err = cursor.All(context.Background(), &books)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var booksInterface []interface{}
		for _, book := range(books) {
			booksInterface = append(booksInterface, book)
		}

		// set content type to JSON
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(utils.GenerateResponse(booksInterface, ""))
	})
}

func saveHandler(client *mongo.Client) http.HandlerFunc {
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

		coll := client.Database(db.Database).Collection(booksCollection)
		_, err = coll.InsertMany(context.TODO(), documents)
		if err != nil {
			utils.JSONError(rw, utils.GenerateResponse(nil, err.Error()), http.StatusInternalServerError)
			return
		}

		var booksInterface []interface{}
		for _, book := range(books) {
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
	database := db.New()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", viewHandler(database)).Methods("GET")
	router.HandleFunc("/books/{id}", viewHandler(database)).Methods("GET")
	router.HandleFunc("/books", saveHandler(database)).Methods("POST")
	return router
}
