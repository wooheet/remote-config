package common

import (
	"fmt"
	"github.com/wooheet/remote-config/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var gGormDB *gorm.DB

func CloseDB() {
	if gGormDB != nil {
		gGormDB.Close()
	}
}

func GetDB() *gorm.DB {
	return gGormDB
}

func Init() {
	db, err := gorm.Open("mysql", dbConnString())
	if err != nil {
		log.Println("[DB Error] ", err)
		panic(err)
	}
	gGormDB = db

	log.Println("[DB] Start DB Migration ... ")
	log.Println("[DB] Start DB ... ")

	if err := db.AutoMigrate(&models.Users{}, &models.Configs{}).Error; err != nil {
		log.Println("[DB Error] ", err)
		panic(err)
	}
}

func dbConnString() (dbConnString string) {
	dbHost := viper.GetString(`database.host`)
	dbHost = "127.0.0.1"
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	dbConnString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName,
	)

	//docker run --platform linux/amd64 -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=users -e MYSQL_USER=test -e MYSQL_PASSWORD=password -d mysql

	return
}
