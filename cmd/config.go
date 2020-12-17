/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/YanJieMao/ToDo/config"
	"github.com/YanJieMao/ToDo/database"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "server config and database config",
	Long: `server config and database config:
	example:
	config --serverPort "9999" --dbType "mysql" --dbHost "localhost" --dbPort "3306" --dbName "todo" --dbParams "parseTime=true&loc=Local" --dbUser "root" --dbPasswd "1997"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		database.Init()
		irisRun()

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&config.ServerPort, "serverPort", "", "", "请输入服务端口")
	configCmd.Flags().StringVarP(&config.DbType, "dbType", "", "", "请输入数据库类型")
	configCmd.Flags().StringVarP(&config.DbHost, "dbHost", "", "", "请输入数据库host地址")
	configCmd.Flags().StringVarP(&config.DbPort, "dbPort", "", "", "请输入数据库端口号")
	configCmd.Flags().StringVarP(&config.DbName, "dbName", "", "", "请输入数据库名称")
	configCmd.Flags().StringVarP(&config.DbParams, "dbParams", "", "", "请输入数据库参数")
	configCmd.Flags().StringVarP(&config.DbUser, "dbUser", "", "", "请输入账号")
	configCmd.Flags().StringVarP(&config.DbPasswd, "dbPasswd", "", "", "请输入密码")

}
