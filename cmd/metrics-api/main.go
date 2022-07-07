package main

import (
	"log"

	common "github.com/hackinggrowth/cafemetrics/metrics-api/internal/commons"
	"github.com/hackinggrowth/cafemetrics/metrics-api/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func viperInit() {
	viper.SetConfigType("json")
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
	common.Init()
	defer common.CloseDB()
	routers.Init(r)

	log.Fatal(r.Run(":8080"))
}
