// Package africastalking provides a simple wrapper for sending SMS via Africa's Talking API.
//
// Example usage:
//
//	package main
//
//	import (
//		"log"
//		"github.com/tech-kenya/africastalkingsms"
//	)
//
//	func main() {
//		apiKey := os.Getenv("atApiKey")
//		username := os.Getenv("atUserName")
//		atShortCode := os.Getenv("atShortCode")
//		sandbox := os.Getenv("sandboxEnv")
//		client, err := africastalking.NewSMSClient(apiKey, username, atShortCode, sandbox)
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
package africastalking // import "github.com/tech-kenya/africastalkingsms"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// NewSMSClient initializes a new Africa's Talking SMS client.
// ensure .env has the below variables
func NewSMSClient(apiKey, username, atShortCode, sandbox string) (*SMSClient, error) {
	// Ensure all required values are set
	if apiKey == "" || username == "" || atShortCode == "" || sandbox == "" {
		return nil, errors.New("missing API credentials: provide API key, username, shortcode and sandbox value")
	}

	return &SMSClient{
		APIKey:     apiKey,
		Username:   username,
		ShortCode:  atShortCode,
		isSandbox:  sandbox,
		HTTPClient: &http.Client{},
	}, nil
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
	// Determine API URL based on environment
	// default is sandbox
	apiURL := "https://api.sandbox.africastalking.com/version1/messaging"
	// production url
	if c.isSandbox == "false" {
		apiURL = "https://api.africastalking.com/version1/messaging"
	}

	// Validate inputs
	if strings.TrimSpace(recipient) == "" || strings.TrimSpace(message) == "" {
		return nil, errors.New("recipient and message cannot be empty")
	}

	// Prepare request data
	data := url.Values{}
	data.Set("username", c.Username)
	data.Set("to", recipient)
	data.Set("message", message)
	data.Set("from", c.ShortCode)

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("apiKey", c.APIKey)

	// Execute HTTP request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %v", err)
	}

	// Debug: Print Raw Response
	fmt.Println("Raw API Response:", string(body))

	// Check if response starts with '{' (indicating JSON)
	if len(body) == 0 || (body[0] != '{' && body[0] != '[') {
		return nil, fmt.Errorf("unexpected API response format: %s", string(body))
	}

	// Parse JSON response
	var smsResponse SMSResponse
	if err := json.Unmarshal(body, &smsResponse); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	// Validate the parsed response
	if len(smsResponse.SMSMessageData.Recipients) == 0 {
		return nil, errors.New("API response is missing recipient details")
	}

	recipientData := smsResponse.SMSMessageData.Recipients[0]

	// Handle unsuccessful SMS sending
	if recipientData.StatusCode != 101 {
		return nil, fmt.Errorf("SMS sending failed for %s: status=%s, statusCode=%d",
			recipientData.Number, recipientData.Status, recipientData.StatusCode)
	}

	return &smsResponse, nil
}
