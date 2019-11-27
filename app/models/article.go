package models

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/volatiletech/sqlboiler/boil"
	"html"
	"minmin/models"
)

//https://developers.google.com/search/docs/data-types/article?hl=vi#amp
type ArticleLogic interface {
	Insert(ctx iris.Context, form *PostForm) error
	GetAll() models.ArticleSlice
}

type articleLogic struct{}

func (articleLogic) GetAll() models.ArticleSlice {
	rs, _ := models.Articles().All(context.Background(), db)

	return rs
}

func (m articleLogic) Insert(ctx iris.Context, form *PostForm) error {
	e := new(models.Article)
	e.Title = form.Title
	e.Content = html.EscapeString(form.Content)

	return e.Insert(context.Background(), db, boil.Infer())
}

func NewArticleLogic() ArticleLogic {
	return &articleLogic{}
}

type PostForm struct {
	Title      string
	Content    string
	Categories []string
}
