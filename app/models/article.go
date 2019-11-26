package models

import (
	"database/sql"
	"fmt"
	uuid "github.com/iris-contrib/go.uuid"
	"minmin/app/conf"
	"time"
)

//https://developers.google.com/search/docs/data-types/article?hl=vi#amp
type Article struct {
	Title       string
	Desc        string
	Author      string
	InLanguage  string
	CatId       uuid.UUID
	Position    string
	Headline    string
	Keywords    string
	ImageThumb  sql.NullString
	ImageS      sql.NullString
	ImageM      sql.NullString
	ImageL      sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	BaseModel
}

const table = "articles"

func (t *Article) Insert() {
	db := conf.DB()
	query := fmt.Sprintf(`INSERT INTO %v`, table)

}
