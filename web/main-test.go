package main

import (
	"fmt"
	"log"
	"todo-mvc/todo/datasource"
)

func mainTest() {

	engine := datasource.NewMysqlEngine()
	sql := "SELECT * FROM user"
	//results, err := engine.Query(sql)
	results, err := engine.QueryInterface(sql)
	//results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal("query", sql, err)
		return
	}
	total := len(results)
	if total == 0 {
		fmt.Println("没有任何数据", sql)
	} else {
		for i, data := range results {
			fmt.Printf("%d = %v\n", i, data)
		}
	}

}
