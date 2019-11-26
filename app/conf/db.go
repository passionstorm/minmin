package conf

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"os"
	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	_db := os.Getenv("DB")
	dbLocal, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+_db)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	boil.SetDB(dbLocal)
	boil.DebugMode = true
	return dbLocal
}
