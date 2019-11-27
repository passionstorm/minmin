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
	GetAll() string
}

type articleLogic struct {
	e models.Article
}

func (m articleLogic) GetAll() string {
	return "hello"
}

func (m articleLogic) Insert(ctx iris.Context, form *PostForm) error {
	m.e.Title = form.Title
	m.e.Content = html.EscapeString(form.Content)

	return m.e.Insert(context.Background(), db, boil.Infer())
}

func NewArticleLogic() ArticleLogic {
	return &articleLogic{}
}

type PostForm struct {
	Title      string
	Content    string
	Categories []string
}
