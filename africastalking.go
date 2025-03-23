// Package africastalking provides a simple wrapper for sending SMS via Africa's Talking API.
//
// Example usage:
//
//	import (
//		"log"
//		"github.com/Tech-Kenya/africastalking-sms-lib"
//	)
//
//	func main() {
//		will load .env variables directly
//		client := africastalking.NewSMSClient()
//		if err != nil {
//			log.Fatal(err)
//		}
//		resp, err := client.SendSMS("+254712345678", "Hello from Africa's Talking!")
//		if err != nil{
//			log.Fatal(err)
//		}
//
//		log.Println(resp)
//	}
package africastalking

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// NewClient initializes a new Africa's Talking SMS client.
func NewSMSClient() (*SMSClient, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("missing API credentials in environment variables")
	}

	apiKey := os.Getenv("atApiKey")
	username := os.Getenv("atUserName")
	atShortCode := os.Getenv("atShortCode")
	if apiKey == "" || username == "" || atShortCode == "" {
		return nil, errors.New("missing API credentials in environment variables")
	}
	return &SMSClient{APIKey: apiKey, Username: username, ShortCode: atShortCode,
		Env: "production", HTTPClient: &http.Client{}}, nil
}

// SendSMS sends an SMS message to the specified recipient.
//
// Example:
//
//	err := client.SendSMS("+254712345678", "Hello!")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// Parameters:
// - recipient: Phone number of the recipient in international format (e.g., +254712345678)
// - message: The text message content
func (c *SMSClient) SendSMS(recipient, message string) (*SMSResponse, error) {
	apiURL := "https://api.sandbox.africastalking.com/version1/messaging"
	data := url.Values{}
	data.Set("username", c.Username)
	data.Set("to", recipient)
	data.Set("message", message)
	data.Set("from", c.ShortCode)

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("apiKey", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var smsResponse SMSResponse
	if err := json.Unmarshal(body, &smsResponse); err != nil {
		return nil, err
	}

	if len(smsResponse.SMSMessageData.Recipients) == 0 || smsResponse.SMSMessageData.Recipients[0].Status != "Success" {
		return nil, errors.New("SMS sending failed, check variables")
	}

	// ensure api response returns code 200
	if smsResponse.SMSMessageData.Recipients[0].StatusCode != 200 {
		return nil, fmt.Errorf("SMS sending failed for %s: status=%s, statusCode=%d",
			smsResponse.SMSMessageData.Recipients[0].Number, smsResponse.SMSMessageData.Recipients[0].Status, smsResponse.SMSMessageData.Recipients[0].StatusCode)
	}

	return &smsResponse, nil
}
