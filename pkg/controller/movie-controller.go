package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daksh-axel/gocrudmovie/pkg/models"
	"github.com/daksh-axel/gocrudmovie/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	mov := models.GetAllMovies()
	res, _ := json.Marshal(mov)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	mov := models.GetMovieByIDDB(id)
	res, _ := json.Marshal(mov)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	newmov := &models.Movie{}
	utils.ParseBody(r, newmov)
	mov := newmov.CreateMovie()
	res, _ := json.Marshal(mov)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 0)
	mov := models.DeleteMovie(id)
	res, _ := json.Marshal(mov)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	updmov := &models.Movie{}
	utils.ParseBody(r, updmov)
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 0)
	currmov := models.GetMovieByIDDB(id)
	updmov.ID = uint(id)
	if updmov.Isbn == "" {
		updmov.Isbn = currmov.Isbn
	}
	if updmov.Name == "" {
		updmov.Name = currmov.Name
	}
	if updmov.Director.Firstname == "" {
		updmov.Director.Firstname = currmov.Director.Firstname
	}
	if updmov.Director.Lastname == "" {
		updmov.Director.Lastname = currmov.Director.Lastname
	}
	updated := models.UpdateMovie(updmov)
	res, _ := json.Marshal(updated)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
