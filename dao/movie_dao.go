package dao

import (
	"log"

	models "github.com/gottsohn/go-fun/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MoviesDAO — movies doa
type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION — collection name
	COLLECTION = "movies"
)

// Connect — create database session connection
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll — fetch all movies
func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
	var movies = make([]models.Movie, 0)
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// FindByID — find movie
func (m *MoviesDAO) FindByID(id string) (models.Movie, error) {
	var movie models.Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert — create new movie
func (m *MoviesDAO) Insert(movie models.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete — drop movie
func (m *MoviesDAO) Delete(movie models.Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// RemoveAll — clear collection
func (m *MoviesDAO) RemoveAll() error {
	_, err := db.C(COLLECTION).RemoveAll(bson.M{})
	return err
}

// Update — update exsting movie
func (m *MoviesDAO) Update(movie models.Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
