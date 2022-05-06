package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wooheet/remote-config/requests"
)

func Config(c *gin.Context) {
	var config requests.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	fmt.Println(config)

	//Extract the access token metadata
	//metadata, err := ExtractTokenMetadata(c.Request)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, "unauthorized")
	//	return
	//}

	//userid, err := FetchAuth(metadata)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, err.Error())
	//	return
	//}

	//config.Token = userid
	//you can proceed to save the Todo to a database
	//but we will just return it to the caller:

	c.JSON(http.StatusCreated, config.Token)
}
