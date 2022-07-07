package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	common "github.com/hackinggrowth/cafemetrics/metrics-api/internal/commons"
	"github.com/hackinggrowth/cafemetrics/metrics-api/models"
	"github.com/hackinggrowth/cafemetrics/metrics-api/requests"
)

func Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, "Retrieve")
}

func Registry(c *gin.Context) {
	var config requests.Configs

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	// TODO: check stored id, tracker type

	metadata, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	userid, err := FetchAuth(metadata)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	log.Println(userid)

	configs := map[string]string{
		"token":        config.Token,
		"tracker_type": config.TrackerType,
		"store_id":     config.StoreId,
	}

	// TODO: insert user id
	common.GetDB().Create(&models.Configs{
		Token:       config.Token,
		TrackerType: config.TrackerType,
		StoreId:     config.StoreId,
		Users: models.Users{
			ID: userid,
		},
	})

	c.JSON(http.StatusCreated, configs)
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, "Update")
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "Delete")
}

func ScriptTag(c *gin.Context) {
	c.JSON(http.StatusOK, "ScriptTag")
}
