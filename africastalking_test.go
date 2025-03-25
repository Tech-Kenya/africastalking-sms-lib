package africastalking // import "github.com/tech-kenya/africastalkingsms"

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
)

// Mock HTTP client for testing
type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// TestSendSMS_Success ensures SMS sending works correctly
func TestSendSMS_Success(t *testing.T) {
	// Mock API response
	mockResponse := `{
		"SMSMessageData": {
			"Message": "Sent to 1/1 Total Cost: KES 0.8000 Message parts: 1",
			"Recipients": [
				{
					"number": "+254712345678",
					"cost": "KES 0.8000",
					"status": "Success",
					"statusCode": 101,
					"messageId": "ATXid_ec7f36df3f46eca883286687796cda82"
				}
			]
		}
	}`

	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		},
	}

	client := SMSClient{
		APIKey:     "test-api-key",
		Username:   "sandbox",
		ShortCode:  "30216",
		HTTPClient: mockClient, // Inject mock client
	}

	resp, err := client.SendSMS("+254712345678", "Hello!")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.SMSMessageData.Recipients[0].Status != "Success" {
		t.Errorf("expected SMS status Success, got %s", resp.SMSMessageData.Recipients[0].Status)
	}
}

// TestSendSMS_Failure ensures SMS failure is handled correctly
func TestSendSMS_Failure(t *testing.T) {
	mockResponse := `{
		"SMSMessageData": {
			"Message": "Failed to send SMS",
			"Recipients": [
				{
					"number": "+254712345678",
					"cost": "KES 0.0000",
					"status": "Failed",
					"statusCode": 400,
					"messageId": ""
				}
			]
		}
	}`

	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		},
	}

	client := SMSClient{
		APIKey:     "test-api-key",
		Username:   "sandbox",
		ShortCode:  "30216",
		HTTPClient: mockClient,
	}

	_, err := client.SendSMS("+254712345678", "Hello!")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

// TestSendSMS_InvalidResponse tests handling of invalid JSON responses
func TestSendSMS_InvalidResponse(t *testing.T) {
	mockResponse := `{ invalid json }`

	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		},
	}

	client := SMSClient{
		APIKey:     "test-api-key",
		Username:   "sandbox",
		ShortCode:  "30216",
		HTTPClient: mockClient,
	}

	_, err := client.SendSMS("+254712345678", "Hello!")
	if err == nil {
		t.Fatalf("expected JSON parse error, got nil")
	}
}

// TestSendSMS_HTTPFailure simulates a network failure
func TestSendSMS_HTTPFailure(t *testing.T) {
	mockClient := &mockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("network failure")
		},
	}

	client := SMSClient{
		APIKey:     "test-api-key",
		Username:   "sandbox",
		ShortCode:  "30216",
		HTTPClient: mockClient,
	}

	_, err := client.SendSMS("+254712345678", "Hello!")
	if err == nil {
		t.Fatalf("expected network failure error, got nil")
	}
}

// TestNewSMSClient_MissingEnv tests missing API credentials
// func TestNewSMSClient_MissingEnv(t *testing.T) {
// 	os.Clearenv() // Clear environment variables

// 	_, err := NewSMSClient()
// 	if err == nil {
// 		t.Fatalf("expected error due to missing env vars, got nil")
// 	}
// }
