package logic

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"html"
	"minmin/app/models"
)

//https://developers.google.com/search/docs/data-types/article?hl=vi#amp
type ArticleLogic interface {
	Insert(ctx iris.Context, form *PostForm) error
	GetAll() models.ArticleSlice
	FindById(id int) (*models.Article, error)
}

type articleLogic struct{}

func (m articleLogic) FindById(id int) (*models.Article, error) {
	return models.Articles(qm.Where("id =?", id)).One(context.Background(), db)
}

func (articleLogic) GetAll() models.ArticleSlice {
	rs, _ := models.Articles().All(context.Background(), db)
	for _, val := range rs {
		val.Content = html.UnescapeString(val.Content)
		val.ImageThumb = null.StringFrom("https://via.placeholder.com/100/ccc")
	}
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