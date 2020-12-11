package main

import (
	"time"
	"todo-mvc/todo/config"
	"todo-mvc/web/controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := newApp()

	//应用App设置
	configation(app)

	//路由设置
	mvcHandle(app)

	config := config.InitConfig()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr), //在端口8080进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
}

//构建App
func newApp() *iris.Application {
	app := iris.New()

	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/", iris.Dir("./public"))

	/* //注册视图文件
	app.RegisterView(iris.HTML("/", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	}) */

	return app
}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	/* 	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
	   		context.JSON(iris.Map{
	   			"errmsg": iris.StatusNotFound,
	   			"msg":    " not found ",
	   			"data":   iris.Map{},
	   		})
	   	})

	   	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
	   		context.JSON(iris.Map{
	   			"errmsg": iris.StatusInternalServerError,
	   			"msg":    " interal error ",
	   			"data":   iris.Map{},
	   		})
	   	}) */
}

//MVC 架构模式处理
func mvcHandle(app *iris.Application) {
	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	//engine := datasource.NewMysqlEngine()

	//用户模块功能
	//userService := service.NewUserService(engine)

	user := mvc.New(app.Party("/user")) //设置路由组
	user.Register(
		//userService,
		sessManager.Start,
	)
	//通过mvc的Handle方法进行控制器的指定
	user.Handle(new(controllers.UserController))
}
