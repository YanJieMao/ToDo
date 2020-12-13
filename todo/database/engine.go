package datasource

import (
	"todo-mvc/todo/datamodles"

	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
)

/**
 * 实例化数据库引擎方法：mysql的数据引擎
 */
func NewMysqlEngine() *xorm.Engine {

	//数据库引擎
	engine, err := xorm.NewEngine("mysql", "root:1997@tcp(127.0.0.1:3306)/todo?charset=utf8")

	//Sync2是Sync的基础上优化的方法
	err = engine.Sync2(
		new(datamodles.User),
	)

	if err != nil {
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}
