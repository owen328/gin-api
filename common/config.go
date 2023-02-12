package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB
var Salt string

func init() {
	initConfig()
	initDB()
}

func initConfig() {
	//workerDir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	fmt.Println("配置初始化")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Salt = viper.GetString("server.salt")
}

func initDB() *gorm.DB {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	loc := viper.GetString("database.loc")
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)
	fmt.Println("数据库连接:", sqlStr)
	db, err := gorm.Open(mysql.Open(sqlStr))
	if err != nil {
		panic(err)
	}
	DB = db
	return DB
}
