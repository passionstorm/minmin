package conf

import (
	"github.com/jmoiron/sqlx"
	"os"
	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

var dbLocal *sqlx.DB

// DB function
func DB() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	_db := os.Getenv("DB")
	dbLocal, _ = sqlx.Connect("mysql", user+":"+password+"@tcp("+host+":3306)/"+_db)
	return dbLocal
}
