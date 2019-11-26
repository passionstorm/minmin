package models

import (
	"context"
	"fmt"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"minmin/models"
)

//https://developers.google.com/search/docs/data-types/article?hl=vi#amp
type Article struct {
	models.Article
}

func (m *Article) InitData() {
	m.Title = null.StringFrom("hello")
	ctx := context.Background()
	err := m.Insert(ctx, db, boil.Infer())
	fmt.Println(err)
}
