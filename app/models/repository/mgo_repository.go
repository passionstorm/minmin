package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MgoRepository interface {
	DB() *mgo.Database
	Save() error
	Update(string, bson.M) error
	Delete(string) error
	FindByID(string) (bson.M, error)
	FindAll() (bson.M, error)
	Collections() ([]string, error)
}

type mgoRepository struct {
	db *mgo.Database
}

func (m mgoRepository) DB() *mgo.Database {
	return m.db
}

func (m mgoRepository) Collections() ([]string, error) {
	return m.db.CollectionNames()
}

func (m mgoRepository) Save() error {
	panic("implement me")
}

func (m mgoRepository) Update(string, bson.M) error {
	panic("implement me")
}

func (m mgoRepository) Delete(string) error {
	panic("implement me")
}

func (m mgoRepository) FindByID(string) (bson.M, error) {
	panic("implement me")
}

func (m mgoRepository) FindAll() (bson.M, error) {
	panic("implement me")
}

func NewBaseRepository(db *mgo.Database) MgoRepository {
	return &mgoRepository{db: db}
}
