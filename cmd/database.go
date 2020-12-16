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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Viper viper global instance
var Viper *viper.Viper

var (
	ServerPort string
	DbType     string
	DbHost     string
	DbPort     string
	DbName     string
	DbParams   string
	DbUser     string
	DbPasswd   string
)

// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database config",
	Long: `Database config. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

/* func configUpdate() {
	fmt.Print("开始更新配置文件")

	Viper = viper.New()
	// scan the file named config in the root directory
	Viper.AddConfigPath("./")
	Viper.SetConfigName("test")

	// read config, if failed, configure by default
	if err := Viper.ReadInConfig(); err == nil {
		log.Println("Read config successfully: ", Viper.ConfigFileUsed())
	} else {
		log.Printf("Read failed: %s \n", err)
		panic(err)
	}

	Viper.ReadInConfig()
	Viper.Set("database.driver", dbType)
	Viper.Set("mysql.dbHost", dbHost)
	Viper.Set("mysql.dbPort", dbPort)
	Viper.Set("mysql.dbName", dbName)
	Viper.Set("mysql.dbParams", dbParams)
	Viper.Set("mysql.dbUser", dbUser)
	Viper.Set("mysql.dbPasswd", dbPasswd)
	fmt.Println(
		Viper.Get("mysql.dbPasswd"),
	)
	//Viper.WriteConfig()
	Viper.WriteConfigAs("test.toml")

} */

func init() {
	rootCmd.AddCommand(databaseCmd)

	databaseCmd.Flags().StringVarP(&ServerPort, "serverPort", "", "", "请输入服务端口")
	databaseCmd.Flags().StringVarP(&DbType, "dbType", "", "", "请输入数据库类型")
	databaseCmd.Flags().StringVarP(&DbHost, "dbHost", "", "", "请输入数据库host地址")
	databaseCmd.Flags().StringVarP(&DbPort, "dbPort", "", "", "请输入数据库端口号")
	databaseCmd.Flags().StringVarP(&DbName, "dbName", "", "", "请输入数据库名称")
	databaseCmd.Flags().StringVarP(&DbParams, "dbParams", "", "", "请输入数据库参数")
	databaseCmd.Flags().StringVarP(&DbUser, "dbUser", "", "", "请输入账号")
	databaseCmd.Flags().StringVarP(&DbPasswd, "dbPasswd", "", "", "请输入密码")

}
