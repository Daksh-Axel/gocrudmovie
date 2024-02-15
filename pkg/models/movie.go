package models

import (
	"github.com/daksh-axel/gocrudmovie/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Movie struct {
	gorm.Model
	Isbn     string    `gorm:""json:"isbn"`
	Name     string    `json:"name"`
	Director *Director `gorm:"embedded"json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func GetAllMovies() []Movie {
	var movies []Movie
	db.Find(&movies)
	return movies
}

func GetMovieByIDDB(id int64) *Movie {
	var movie Movie
	db.Where("ID=?", id).Find(&movie)
	return &movie
}

func (m *Movie) CreateMovie() *Movie {
	db.Create(&m)
	return m
}

func DeleteMovie(id int64) *Movie {
	var mov Movie
	db.Where("ID=?", id).Delete(&mov)
	return &mov
}

func UpdateMovie(m *Movie) *Movie {
	db.Save(&m)
	return m
}
