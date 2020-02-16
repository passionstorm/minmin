package db

import (
	"context"
	"fmt"
	"github.com/ztrue/tracerr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"minmin/app/utils/m"
	"time"
)

const (
	hosts    = "18.139.226.39:27017"
	database = "mintoot"
	username = "admin"
	password = "qwer1234"
)

func NewClient() *mongo.Client {
	options := options.Client().ApplyURI(fmt.Sprintf(`mongodb://%v:%v@%v/%v`, username, password, hosts, database))
	client, err := mongo.Connect(context.TODO(), options)
	m.HandlErr(tracerr.Wrap(err))
	err = client.Ping(context.TODO(), nil)
	m.HandlErr(tracerr.Wrap(err))
	return client
}

func DefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

var client *mongo.Client

func GetDB() *mongo.Database {
	client = NewClient()
	//defer client.Disconnect(DefaultContext())
	db := client.Database(database)

	return db
}

func Disconnect() {
	if client != nil {
		client.Disconnect(context.Background())
	}
}
