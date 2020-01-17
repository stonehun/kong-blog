package boot

import (
	"github.com/kataras/iris/v12"
	"kong_blog/configs"
	"kong_blog/middleware"
	"kong_blog/routers"
)

func Run() {
	//获取配置
	config := configs.GetConfig()

	//获取服务对象
	app := iris.New()

	//开发环境设置
	if config.IsDev() {
		app.Use(middleware.LogRequest)
		app.Done(middleware.LogResponse)
		//设置日志级别
		app.Logger().SetLevel("debug")
	}

	//注册路由
	routers.RegisterRouter(app)

	//应用启动
	_ = app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":8080"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
}
