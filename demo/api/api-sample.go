package api

import (
	"log"
	"net/http"
	"github.com/Tech-Kenya/africastalking-sms-lib"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := 
	r := gin.Default()
	r.POST("/send-sms", sendSMS)
	log.Println("server running on http://localhost:8080")
	r.Run(":8080")
}

func sendSMS(c *gin.Context) {
	type body struct {
		Recipient string `json:"recipient" binding:"required"`
		Message   string `json:"message" binding:"required"`
	}
	var reqBody body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
}
