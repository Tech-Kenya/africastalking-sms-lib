## Africastalking Golang SMS library

This is a Golang library for sending SMS messages using the Africa's Talking API. It provides a simple interface to send messages, check message status, and manage contacts.

## Installation

To install the library, use the following command:

```bash
go get github.com/Tech-Kenya/africastalking_sms_go
```

## To get started locally

1. Clone the repository:

```bash
git clone https://github.com/Tech-Kenya/africastalking_sms_go.git
```

2. cd into the project directory

```bash
africastalking_sms_go
```

3. Copy the `.env.example` file to `.env` and fill in your Africa's Talking credentials:

```bash
cp .env.example .env
```

`Ensure you have Golang 1,18+ installed on your machine and you have an API key from Africa's Talking.`

- Shortcode or Sender ID: <https://account.africastalking.com/apps/sandbox/sms/shortcodes/create>
- API Key: <https://account.africastalking.com/apps/sandbox/settings/key>

4. Install the dependencies:

```bash
go mod tidy
```

5. Run the example:

```bash
go run demo/cli.go
go run demo/api.go
```
