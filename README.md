
# Telegram Go SDK

This project is a Go SDK for interacting with the Telegram Bot API. With this SDK, you can easily manage your Telegram bot and utilize various features of the Telegram API.

## Features
- Send messages to users
- Manage Callback Queries
- Handle user accounts and related information
- Set up and manage Webhooks for receiving bot events

## Installation

To install this SDK, use the following command:

```bash
go get github.com/parparvaz/telegram-golang-sdk
```

This command will install the library in your Go project.

## Usage

###  . Simple Message Sending

```go
package main

import (
	"context"
	"github.com/parparvaz/telegram-golang-sdk"
	"log"
)

func main() {
	client := telegram.NewClient("YOUR_BOT_TOKEN", "YOUR_SECRET")
	chatID := "CHAT_ID"
	message := "Hello! This is a test message."

	res, err := client.NewSendMessageService().
		Text(chatID).
		ChatID(message).
		Do(context.Background())
	if err != nil {
		log.Println(err)
		panic(1)
	}
	log.Println(res)
	
}
```

###  . Setting Up a Webhook

```go
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/parparvaz/telegram-golang-sdk"
)

func main() {
	client := telegram.NewClient("YOUR_BOT_TOKEN", "")

	r := gin.Default()
	r.POST("/webhook", func(c *gin.Context) {
		c.Status(200)
	})
	client.NewSetWebhookService().
		URL("https://yourdomain.com/webhook").
		SecretToken("YOUR_SECRET").
		Do(context.Background())
	r.Run(":8080")
}
```

## Project Structure
The project consists of various files, each responsible for handling different operations:

- **account_service.go**: Manages user account information
- **callback_query_service.go**: Handles Callback Queries
- **send_service.go**: Sends messages and media
- **webhook_service.go**: Sets up and manages Webhooks

## Contributing
If you're interested in improving this project, feel free to submit Pull Requests or report issues. All contributions are welcome.

## License
This project is licensed under the MIT License. For more details, refer to the [LICENSE](LICENSE) file.
