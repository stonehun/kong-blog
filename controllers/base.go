package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"kong_blog/models"
)

type BaseController struct {
	authUser models.Auth
}

//获取授权信息
func (c *BaseController) AuthUser(ctx iris.Context) models.Auth {
	if c.authUser.Id == 0 {
		user := ctx.Values().Get("jwt").(*jwt.Token)
		userMap := user.Claims.(jwt.MapClaims)
		fmt.Println("11111111111", userMap)
		c.authUser = models.Auth{
			Id:       uint(userMap["id"].(float64)),
			Name:     userMap["name"].(string),
			NickName: userMap["nick_name"].(string),
		}
	}

	return c.authUser
}

//成功响应
func (c *BaseController) SuccessResponse(ctx iris.Context, data interface{}) {
	ctx.Values().Set("body", struct {
		Code    int64       `json:"code"`
		Message string      `json:"string"`
		Data    interface{} `json:"data"`
	}{
		Code:    1,
		Message: "请求成功",
		Data:    data,
	})
}

//失败响应
func (c *BaseController) FailResponse(ctx iris.Context, message string) {
	ctx.Values().Set("body", struct {
		Code    int64  `json:"code"`
		Message string `json:"string"`
	}{
		Code:    0,
		Message: message,
	})
}
