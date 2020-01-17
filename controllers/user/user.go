package user

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type Response struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func Add(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)

	fmt.Println(user.Claims)

	ctx.Values().Set("body", Response{
		Name: "yangzai",
		Code: "1233",
	})
}
