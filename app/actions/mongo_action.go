package actions

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
	"minmin/app/models/repository"
	"minmin/app/utils/mm"
	"net/http"
	"time"
)

type mongoAction struct {
	repository.MgoRepository
}

func (m mongoAction) Routes(c *echo.Echo) {
	c.Use(middleware.CORS())
	g := c.Group("db")
	g.GET("", m.listCollection)
	g.GET("/", m.listCollection)
	g.GET("/:collect", m.fetchAll)
	g.GET("/:collect/", m.fetchAll)
	g.POST("/:collect/new", m.addNew)
	g.GET("/:collect/:id", m.fetchId)
	g.GET("/:collect/:id/", m.fetchId)
	g.DELETE("/:collect/:id", m.deleteById)
	//g.POST("/:collect/:id/edit", m.edit)

}

func (m mongoAction) deleteById(c echo.Context) error {
	db := m.DB()
	collectStr := c.Param("collect")
	idStr := c.Param("id")
	if err := db.C(collectStr).RemoveId(bson.ObjectIdHex(idStr)); err != nil {
		return c.JSON(http.StatusBadRequest, mm.NotFound())
	}
	return c.JSON(http.StatusOK, mm.Map{"status": true})
}

func (m mongoAction) addNew(c echo.Context) error {
	db := m.DB()
	collectStr := c.Param("collect")
	frm := mm.Map{}
	err := c.Bind(&frm)
	mm.HandlErr(err)
	slugStr := ""
	if val, ok := frm["name"]; ok {
		slugStr = slug.Make(val.(string))
	}
	id := bson.NewObjectId()
	idStr := id.Hex()
	if slugStr != "" {
		frm["slug"] = fmt.Sprintf(`%v.%v`, slugStr, idStr)
	} else {
		frm["slug"] = ""
	}
	now := time.Now()

	frm["_id"] = id
	frm["created_at"] = now.String()
	frm["updated_at"] = now.String()
	if err = db.C(collectStr).Insert(frm); err != nil {
		return c.JSON(http.StatusBadRequest, mm.NewError(err))
	}
	frm["_id"] = idStr
	log.Printf("[created] %v", frm)
	return c.JSON(http.StatusCreated, frm)
}

func (m mongoAction) fetchId(c echo.Context) error {
	db := m.DB()
	collectStr := c.Param("collect")
	idStr := c.Param("id")
	rs := mm.Map{}
	q := db.C(collectStr).FindId(bson.ObjectIdHex(idStr))
	if err := q.One(&rs); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, rs)
}

func (m mongoAction) fetchAll(c echo.Context) error {
	db := m.DB()
	collectParam := c.Param("collect")
	rs := []mm.Map{}
	q := db.C(collectParam).Find(nil)
	names, _ := db.CollectionNames()
	if _, ok := mm.FindStr(names, collectParam); ok == false {
		return c.JSON(http.StatusNotFound, mm.NotFound())
	}
	if err := q.All(&rs); err != nil {
		return c.JSON(http.StatusNotFound, mm.NewError(err))
	}
	return c.JSON(http.StatusOK, rs)
}

func (m mongoAction) listCollection(c echo.Context) error {
	db := m.DB()
	rs, _ := db.CollectionNames()
	return c.JSON(http.StatusOK, rs)
}

func NewMongoAction(r repository.MgoRepository) BaseAction {
	return &mongoAction{r}
}
