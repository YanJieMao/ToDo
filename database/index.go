package database

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/YanJieMao/ToDo/config"
	"github.com/YanJieMao/ToDo/model/pojo"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
)

var once sync.Once

// DB 数据库连接实例
var DB *xorm.Engine

func Init() {

	fmt.Println(config.DbType)

	once.Do(func() {
		dbType := config.DbType

		switch dbType {
		case "mysql":
			initMysql()
		default:
			panic(errors.New("only support mysql"))
		}

		// 顺序不能错，否则生成的表不能按照配置的规则命名
		configDB()
		initTable()
	})

}

// 初始化，当使用的数据库为Mysql时
func initMysql() {
	/* //dbType := config.DbType
	dbHost := config.DbHost
	dbPort := config.DbPort
	dbName := config.DbName
	dbParams := config.DbParams
	dbUser := config.DbUser
	dbPasswd := config.DbPasswd */

	//dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbUser, dbPasswd, dbHost, dbPort, dbName, dbParams)
	dbURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", config.DbUser, config.DbPasswd, config.DbHost, config.DbPort, config.DbName, config.DbParams)
	fmt.Println(dbURL)

	var err error
	DB, err = xorm.NewEngine("mysql", dbURL)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
}

// 自动同步表结构，如果不存在则创建
func initTable() {
	// 自动创建表
	err := DB.Sync2(new(pojo.User), new(pojo.ToDoList))
	if err != nil {
		log.Printf("同步数据库和结构体字段失败:%v\n", err)
		panic(err)
	}
}

// 设置可选配置
func configDB() {
	// 设置日志等级，设置显示sql，设置显示执行时间
	DB.SetLogLevel(xorm.DEFAULT_LOG_LEVEL)
	DB.ShowSQL(true)
	DB.ShowExecTime(true)

	// 指定结构体字段到数据库字段的转换器
	// 默认为core.SnakeMapper
	// 但是我们通常在struct中使用"ID"
	// 而SnakeMapper将"ID"转换为"i_d"
	// 因此我们需要手动指定转换器为core.GonicMapper{}
	DB.SetMapper(core.GonicMapper{})
}
