package africastalking

// SMSClient holds API credentials
type SMSClient struct {
	APIKey    string
	Username  string
	ShortCode string
	Env       string
}

// Response struct for JSON parsing
type SMSResponse struct {
	SMSMessageData struct {
		Message    string `json:"Message"`
		Recipients []struct {
			Number     string `json:"number"`
			Cost       string `json:"cost"`
			Status     string `json:"status"`
			StatusCode int    `json:"statusCode"`
			MessageID  string `json:"messageId"`
		} `json:"Recipients"`
	} `json:"SMSMessageData"`
}

// Request payload struct
type SMSRequest struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}
