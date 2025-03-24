[![Go Report Card](https://goreportcard.com/badge/github.com/Tech-Kenya/africastalking-sms-lib)](https://goreportcard.com/report/github.com/Tech-Kenya/africastalking-sms-lib)
![Go Version](https://img.shields.io/badge/Go-1.24-blue?style=flat&logo=go)
[![Go Reference](https://pkg.go.dev/badge/github.com/Tech-Kenya/africastalking-sms-lib.svg)](https://pkg.go.dev/github.com/Tech-Kenya/africastalking-sms-lib)

## Africastalking Golang SMS library

This is a Golang library for sending SMS messages using the Africa's Talking API. It provides a simple interface to send messages, check message status, and manage contacts.

## Installation

To install the library, use the following command:

```bash
go get github.com/Tech-Kenya/africastalking-sms-lib@latest
```

## To get started locally

1. Clone the repository:

```bash
git clone https://github.com/Tech-Kenya/africastalking-sms-lib.git
```

2. cd into the project directory

```bash
cd africastalking-sms-lib/demo
```

3. Copy the `.env.example` file to `.env` and fill in your Africa's Talking credentials:

```bash
cp .env.example .env
cp .env.example .envrc
```

`Ensure you have Golang 1,18+ installed on your machine and you have an API key from Africa's Talking.`

- Shortcode or Sender ID: <https://account.africastalking.com/apps/sandbox/sms/shortcodes/create>

###### note that after generating a new API key, it might take upto 15 mins for it to become active

- API Key: <https://account.africastalking.com/apps/sandbox/settings/key>

4. Install the dependencies:

```bash
go mod tidy
```

5. Run the example:

```bash
direnv allow # if you are using direnv
# or
export $(cat .env | xargs) # if you are not using direnv
#then
go run .
```

## Usage


see demo/api for a simple API example
```bash
go mod tidy #download dependencies
```

```bash
go run . #start the api server
```

POST Request


```bash
curl -X POST http://localhost:8080/send-sms \
     -H "Content-Type: application/json" \
     -d '{
           "recepient": "Your number",
           "message": "Hello from Gin API"
         }'
```

Response
```json
{"Message":"Sent to 1/1 Total Cost: KES 0.8000 Message parts: 1","Recipients":[{"number":"+254....","cost":"KES 0.8000","status":"Success","statusCode":101,"messageId":"ATXid_ad8a62b0680a41351b1ea383b9b66fd1"}]}
```

![sample](demo/image.png)

## Known bugs

```json
2025/03/24 11:35:51 Failed to send SMS:failed to parse API response: invalid character 'T' looking for beginning of value
```