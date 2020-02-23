package db

import (
	"github.com/ztrue/tracerr"
	mgo "gopkg.in/mgo.v2"
	"time"
)

const (
	hosts    = "18.139.226.39:27017"
	database = "mintoot"
	username = "admin"
	password = "qwer1234"
)

func GetMongoDB() (*mgo.Database, error) {
	dialInfo, err := mgo.ParseURL(hosts)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	dialInfo.Timeout = 5 * time.Second
	dialInfo.Username = username
	dialInfo.Password = password
	dialInfo.Database = database
	dialInfo.Mechanism = ""

	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	db := sess.DB(database)
	return db, nil
}
