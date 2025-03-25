package africastalking // import "github.com/tech-kenya/africastalking-sms"

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SMSClient holds API credentials
type SMSClient struct {
	APIKey     string
	Username   string
	ShortCode  string
	isSandbox  string
	HTTPClient HTTPClient
}

// Response struct for JSON parsing
type SMSResponse struct {
	SMSMessageData struct {
		Message    string `json:"message"`
		Recipients []struct {
			Number     string `json:"number"`
			Cost       string `json:"cost"`
			Status     string `json:"status"`
			StatusCode int    `json:"statusCode"`
			MessageID  string `json:"messageId"`
		} `json:"recipients"`
	} `json:"SMSMessageData"`
}

// Request payload struct
type SMSRequest struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}
