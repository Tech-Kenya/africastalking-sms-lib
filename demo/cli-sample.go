package main

import (
	"fmt"
	"log"

	"github.com/Tech-Kenya/africastalking-sms-lib"
)

// go run .
func main() {
	client, err := africastalking.NewSMSClient()
	if err != nil {
		log.Fatal(err)
	}
	// change the phone number
	response, err := client.SendSMS("+254...", "Hello from Go!")
	if err != nil {
		log.Fatal("Failed to send SMS:", err)
	}

	fmt.Println("Response:", response)
}
