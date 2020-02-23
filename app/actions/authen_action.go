package actions

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
	"minmin/app/utils/mm"
	"net/http"
)

type AuthenAction struct {
}


func (AuthenAction) Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		fmt.Println(e.Routes())

		return c.String(http.StatusOK, "ok")
	})
	e.GET("/login", postLogin)
	//e.POST("/register", postRegister)
}


//func getInfo(c echo.Context) error {
//	tokenStr := c.ghh("Authorization")
//	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method")
//		}
//		return []byte("secret"), nil
//	})
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok || !token.Valid {
//		m.HandlErr(tracerr.Wrap(err))
//	}
//
//	return claims
//}

func postLogin(c echo.Context) error {
	defer mm.CoverErr(c)
	//db := db.GetDB()
	//col := db.Collection("users")
	_, err := c.FormParams()
	mm.HandlErr(tracerr.Wrap(err))
	byt := []byte(`"num":6.13,"strs":["a","b"]}`)
	var dat mm.Map
	err = json.Unmarshal(byt, dat)
	mm.HandlErr(tracerr.Wrap(err))
	return nil
	//var result m.Map
	//err = col.FindOne(context.TODO(), bson.D{{"username", form["username"]}}).Decode(&result)
	//m.HandlErr(tracerr.Wrap(err), "Invalid username")
	//
	//err = bcrypt.CompareHashAndPassword([]byte(result["password"].(string)), []byte(form.Get("password").(string)))
	//m.HandlErr(tracerr.Wrap(err), "Invalid password")
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"username":  result["username"],
	//	"firstname": result["firstname"],
	//	"lastname":  result["lastname"],
	//})
	//tokenStr, err := token.SignedString([]byte("secret"))
	//m.HandlErr(tracerr.Wrap(err), "Error while generating token,Try again")
	//
	//ctx.StatusCode(200)
	//return m.Map{"token": tokenStr}

}

//func postRegister(c echo.Context) error {
//	defer m.CoverErr(ctx)
//	var form m.Map
//	db := db.GetDB()
//	col := db.Collection("users")
//	err := ctx.ReadJSON(&form)
//	m.HandlErr(tracerr.Wrap(err))
//
//	var result m.Map
//	err = col.FindOne(context.TODO(), bson.D{{"username", form["username"]}}).Decode(&result)
//
//	var inserted *mongo.InsertOneResult
//	if m.IsEmptyMgo(err) {
//		hash, err := bcrypt.GenerateFromPassword([]byte(form["password"].(string)), 5)
//		m.HandlErr(tracerr.Wrap(err))
//		form["password"] = string(hash)
//
//		inserted, err = col.InsertOne(context.TODO(), form)
//		m.HandlErr(tracerr.Wrap(err))
//	} else {
//		m.HandlErr(errors.New("username is exists"))
//	}
//	ctx.StatusCode(201)
//	return m.Map{"message": "user created successfully", "id": inserted.InsertedID}
//}
