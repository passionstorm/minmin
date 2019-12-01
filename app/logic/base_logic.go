package logic

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // must import
	"github.com/volatiletech/sqlboiler/boil"
	"os"
	"time"
)

type BaseModel struct {
}

var db *sql.DB

func Load() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	var err error
	dsn := user + ":" + pass + "@tcp(" + host + ":3306)/" + dbname + "?parseTime=true&loc=Asia%2FTokyo"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic("Can't connect to " + dsn)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	//defer db.Close()
	boil.SetDB(db)
	boil.DebugMode = true
}
