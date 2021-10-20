package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/library/routers"
)

const (
	port string = ":8080"
)

func main() {

	// creating a mux router (which will serve as a primary router)
	router := mux.NewRouter().StrictSlash(true)
	// mount entity specific routers
	mount(router, "/", routers.BookRouter())

	fmt.Printf("Listening on port: %s\n", port[1:])
	log.Fatal(http.ListenAndServe(port, router))
}

// mount adds the entity specific routers to primary router
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
