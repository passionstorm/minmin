package main

import (
	"minmin/app/actions"
	"minmin/app/bootstrap"
	"minmin/app/helpers/db"
	"minmin/app/models/repository"
	"minmin/app/utils/mm"
)

func init() {
	app := bootstrap.NewApp()
	db, err := db.GetMongoDB()
	mm.HandlErr(err)
	rp := repository.NewBaseRepository(db)
	app.ImportRoutes(
		actions.AuthenAction{},
		actions.NewMongoAction(rp),
	)
	app.Debug = true
	app.Listen(":9999")
}

func main() {

}
