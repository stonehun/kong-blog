package middleware

import (
	"errors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"kong_blog/configs"
	"kong_blog/models"
)

//请求响应中间件
func LogRequest(ctx iris.Context) {

	method := ctx.Request().Method
	path := ctx.Request().URL.String()
	body, _ := ctx.GetBody()

	ctx.Application().Logger().Debugf("请求地址:%s 请求方法:%s 请求内容:\n%s", path, method, body)

	ctx.Next()
}

//打印出参
func LogResponse(ctx iris.Context) {
	//res:= ctx.Values()
	//fmt.Println("1111111111111111")
	//ctx.Application().Logger().Debugf("响应内容:%s", res)
	//ctx.Next()
}

//jwt配置
func Auth() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.GetConfig().JwtKey), nil
		},
		ErrorHandler: func(context context.Context, e error) {
			_, err := context.JSON(models.NewResult(errors.New("授权失败"), 0))
			if err != nil {
				context.Application().Logger().Warn(err)
				return
			}
		},

		Extractor:     jwt.FromAuthHeader,
		SigningMethod: jwt.SigningMethodHS256,
	})
}
