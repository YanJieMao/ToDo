package main

import (
	"github.com/YanJieMao/ToDo/todo/config"
	"github.com/YanJieMao/ToDo/todo/middleware"
	"github.com/YanJieMao/ToDo/todo/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	// Set logger level
	app.Logger().SetLevel(config.Viper.GetString("server.logger.level"))
	// Add recover to recover from any http-relative panics
	app.Use(recover.New())
	// Add logger to log the requests to the terminal
	app.Use(logger.New())

	// Globally allow options method to enable CORS
	app.AllowMethods(iris.MethodOptions)
	// Add global CORS handler
	app.Use(middleware.CORS)

	// Router
	route.Route(app)

	// Listen in 8888 port
	app.Run(iris.Addr(config.Viper.GetString("server.addr")), iris.WithoutServerError(iris.ErrServerClosed))
}
