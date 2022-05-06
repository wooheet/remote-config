package handler

import (
	"github.com/wooheet/remote-config/common"
	"github.com/wooheet/remote-config/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wooheet/remote-config/requests"
)

func Config(c *gin.Context) {
	var config requests.Configs

	// TODO: http type별 분기

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	// TODO: check stored id, tracker type

	common.GetDB().Create(&models.Configs{
		Token:       config.Token,
		TrackerType: config.TrackerType,
		StoreId:     config.StoreId,
	})

	// TODO: token metadata도 저장을 할것인가?
	//Extract the access token metadata
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
		"access_uuid":  metadata.AccessUuid,
		"user_id":      metadata.UserId,
		"tracker_type": config.TrackerType,
		"store_id":     config.StoreId,
	}

	c.JSON(http.StatusCreated, configs)
}
