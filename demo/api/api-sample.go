package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Tech-Kenya/africastalking-sms-lib"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Client *africastalking.SMSClient
}

func main() {
	client := &africastalking.SMSClient{
		APIKey:    os.Getenv("atApiKey"),
		ShortCode: os.Getenv("atShortCode"),
		Username:  os.Getenv("atUserName"),
	}

	handler := Handler{Client: client}

	r := gin.Default()
	r.POST("/send-sms", handler.sendSMS)

	log.Println("server running on http://localhost:8080")
	r.Run(":8080")
}

func (h *Handler) sendSMS(c *gin.Context) {
	type body struct {
		Recepient string `json:"recepient" binding:"required"`
		Message   string `json:"message" binding:"required"`
	}
	var reqBody body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := h.Client.SendSMS(reqBody.Recepient, reqBody.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.SMSMessageData)
}
