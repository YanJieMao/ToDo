package cmd

import (
	"fmt"

	"github.com/YanJieMao/ToDo/config"
	"github.com/YanJieMao/ToDo/middleware"
	"github.com/YanJieMao/ToDo/route"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func irisRun() {
	//database.InitDB()
	app := iris.New()
	// Set logger level
	app.Logger().SetLevel("debug")
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
	fmt.Println(config.ServerPort)
	// Listen  port
	app.Run(iris.Addr(":"+config.ServerPort), iris.WithoutServerError(iris.ErrServerClosed))

}
