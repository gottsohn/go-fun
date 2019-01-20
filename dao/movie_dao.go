package dao

import (
	"log"

	models "github.com/gottsohn/go-fun/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
	var movies []models.Movie = make([]models.Movie, 0)
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesDAO) FindById(id string) (models.Movie, error) {
	var movie models.Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesDAO) Insert(movie models.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesDAO) Delete(movie models.Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *MoviesDAO) Update(movie models.Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
