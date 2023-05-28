package main

import (
	"example/web-service-gin/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRate(c *gin.Context) {
	bitcoinPrice, err := services.GetBitcoinPrice()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, bitcoinPrice)
}

func PostSubscribe(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		return
	}

	err := services.SubscribeEmail(email)
	if err != nil {
		log.Print(err)
		if err.Error() == "email exist" {
			c.JSON(http.StatusConflict, "")
		} else {
			c.JSON(http.StatusInternalServerError, "")
		}
	}
}

func PostSendEmails(c *gin.Context) {
	err := services.SendEmails()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}
