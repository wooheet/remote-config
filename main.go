package main

import (
	"github.com/wooheet/remote-config/common"
	"github.com/wooheet/remote-config/models"
	"github.com/wooheet/remote-config/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func viperInit() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	viperInit()
	common.RedisInit()
	models.Init()
	defer models.CloseDB()
	routers.Init(r)

	log.Fatal(r.Run(":8080"))
}
