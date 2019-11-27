package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // must import
	"github.com/volatiletech/sqlboiler/boil"
)

type BaseModel struct {
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp([localhost]:3306)/minmin?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	boil.SetDB(db)
	boil.DebugMode = true
}
