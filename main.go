package main

import (
	"github.com/YanJieMao/ToDo/cmd"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cmd.Execute()

}
