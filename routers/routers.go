package routers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"kong_blog/controllers"
	"kong_blog/middleware"
)

func RegisterRouter(application *iris.Application) {
	//登录
	mvc.New(application.Party("/api")).Handle(controllers.NewAuthController())

	//需要授权api
	apiRouter := application.Party("/api", middleware.Auth().Serve)

	//用户管理
	mvc.New(apiRouter.Party("/user")).Handle(controllers.NewUserController())

}
