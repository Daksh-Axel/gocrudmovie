package routes

import (
	"github.com/daksh-axel/gocrudmovie/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(r *mux.Router) {
	r.HandleFunc("/movie", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", controller.GetMovieByID).Methods("GET")
	r.HandleFunc("/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
}
