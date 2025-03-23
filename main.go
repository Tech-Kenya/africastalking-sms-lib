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
//		client := africastalking.NewClient("sandbox", "your_api_key")
//		err := client.SendSMS("+254712345678", "Hello from Africa's Talking!", "30216")
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package africastalking

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// NewClient initializes a new Africa's Talking SMS client.
func NewSMSClient() (*SMSClient, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("atApiKey")
	username := os.Getenv("atUserName")
	atShortCode := os.Getenv("atShortCode")
	if apiKey == "" || username == "" || atShortCode == "" {
		return nil, errors.New("missing API credentials in environment variables")
	}
	return &SMSClient{APIKey: apiKey, Username: username, ShortCode: atShortCode, Env: "production"}, nil
}

// SendSMS sends an SMS message to the specified recipient.
//
// Example:
//
//	err := client.SendSMS("+254712345678", "Hello!", "30216")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// Parameters:
// - recipient: Phone number of the recipient in international format (e.g., +254712345678)
// - message: The text message content
// - shortcode: The shortcode or sender ID used for sending
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

	client := &http.Client{}
	resp, err := client.Do(req)
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

	return &smsResponse, nil
}
