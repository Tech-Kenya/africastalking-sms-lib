package main

import (
	"fmt"
	"log"

	"github.com/Tech-Kenya/africastalking-sms-lib"
)

func main() {
	client, err := africastalking.NewSMSClient()
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.SendSMS("*************", "Hello from Go!")
	if err != nil {
		log.Fatal("Failed to send SMS:", err)
	}

	fmt.Println("Response:", response)
}
