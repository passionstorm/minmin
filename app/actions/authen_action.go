package actions

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/ztrue/tracerr"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"minmin/app/pkg/mgo/db"
	"minmin/app/utils/m"
)

type AuthenAction struct {
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func (AuthenAction) PostLogin(ctx iris.Context) interface{} {
	defer m.CoverErr(ctx)
	var form m.Map
	db := db.GetDB()
	col := db.Collection("users")
	err := ctx.ReadJSON(&form)
	m.HandlErr(tracerr.Wrap(err))

	var result m.Map

	err = col.FindOne(context.TODO(), bson.D{{"username", form["username"]}}).Decode(&result)
	if err != nil{
		m.HandlErr(tracerr.New("Invalid username"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(result["password"].(string)), []byte(form["password"].(string)))
	m.HandlErr(tracerr.New("Invalid password"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result["username"],
		"firstname": result["firstname"],
		"lastname":  result["lastname"],
	})
	tokenStr, err := token.SignedString([]byte("secret"))
	m.HandlErr(tracerr.New("Error while generating token,Try again"))
	ctx.StatusCode(200)
	return m.Map{"token": tokenStr}

}

func (c *AuthenAction) PostRegister(ctx iris.Context) interface{} {
	defer m.CoverErr(ctx)
	var form m.Map
	db := db.GetDB()
	col := db.Collection("users")
	err := ctx.ReadJSON(&form)
	m.HandlErr(tracerr.Wrap(err))

	var result m.Map
	err = col.FindOne(context.TODO(), bson.D{{"username", form["username"]}}).Decode(&result)

	if m.IsEmptyMgo(err) {
		hash, err := bcrypt.GenerateFromPassword([]byte(form["password"].(string)), 5)
		m.HandlErr(tracerr.Wrap(err))
		form["password"] = string(hash)

		_, err = col.InsertOne(context.TODO(), form)
		m.HandlErr(tracerr.Wrap(err))
	} else {
		m.HandlErr(errors.New("username is exists"))
	}
	ctx.StatusCode(201)
	return m.Map{"message": "user created successfully", "success": 1}
}
