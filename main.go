package main

import (
	"log"

	"github.com/wooheet/remote-config/common"
	"github.com/wooheet/remote-config/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func viperInit() {
	viper.SetConfigFile(`./config.json`)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	viperInit()
	common.RedisInit()
	common.Init()
	defer common.CloseDB()
	routers.Init(r)

	log.Fatal(r.Run(":8080"))
}
