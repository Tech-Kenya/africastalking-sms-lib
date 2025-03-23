package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tech-Kenya/africastalking-sms-lib"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Client *africastalking.SMSClient
}

func main() {
	client, err := africastalking.NewSMSClient()

	if err != nil {
		log.Fatal(err)
	}

	handler := Handler{Client: client}

	r := gin.Default()
	r.POST("/send-sms", handler.sendSMS)

	log.Println("server running on http://localhost:8080")
	r.Run(":8080")
}

// sendSMS post request endpoint
func (h *Handler) sendSMS(c *gin.Context) {
	type body struct {
		Recepient string `json:"recepient" binding:"required"`
		Message   string `json:"message" binding:"required"`
	}
	var reqBody body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	resp, err := h.Client.SendSMS(reqBody.Recepient, reqBody.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send SMS", "details": err.Error(),
		})
		return
	}

	for _, recipient := range resp.SMSMessageData.Recipients {
		if recipient.StatusCode != 101 && recipient.StatusCode != 200 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to send SMS",
				"details": fmt.Sprintf("SMS sending failed for %s: status=%s, statusCode=%d",
					recipient.Number, recipient.Status, recipient.StatusCode),
			})
			return
		}
	}

	c.JSON(http.StatusOK, resp.SMSMessageData)
}
