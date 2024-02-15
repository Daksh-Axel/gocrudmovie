package main

import (
	"log"
	"net/http"

	"github.com/daksh-axel/gocrudmovie/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
