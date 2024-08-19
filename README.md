
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

### 1. Simple Message Sending
```go
package main

import (
    "github.com/parparvaz/telegram-golang-sdk"
)

func main() {
    client := telegram.NewClient("YOUR_BOT_TOKEN")
    chatID := "CHAT_ID"
    message := "Hello! This is a test message."
    
    client.SendMessage(chatID, message)
}
```

### 2. Sending Message with Inline Buttons
```go
package main

import (
    "github.com/parparvaz/telegram-golang-sdk"
)

func main() {
    client := telegram.NewClient("YOUR_BOT_TOKEN")
    chatID := "CHAT_ID"
    
    buttons := [][]telegram.InlineKeyboardButton{
        {
            {Text: "Button 1", CallbackData: "callback_1"},
            {Text: "Button 2", CallbackData: "callback_2"},
        },
    }
    keyboard := telegram.InlineKeyboardMarkup{InlineKeyboard: buttons}
    
    client.SendMessageWithKeyboard(chatID, "Please choose one of the buttons:", keyboard)
}
```

### 3. Sending a Photo
```go
package main

import (
    "github.com/parparvaz/telegram-golang-sdk"
)

func main() {
    client := telegram.NewClient("YOUR_BOT_TOKEN")
    chatID := "CHAT_ID"
    photoURL := "https://example.com/photo.jpg"
    
    client.SendPhoto(chatID, photoURL, "This is a test photo.")
}
```

### 4. Setting Up a Webhook
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/parparvaz/telegram-golang-sdk"
)

func main() {
    client := telegram.NewClient("YOUR_BOT_TOKEN")
    r := gin.Default()
    
    r.POST("/webhook", func(c *gin.Context) {
        update := telegram.Update{}
        if err := c.ShouldBindJSON(&update); err == nil {
            if update.Message != nil {
                chatID := update.Message.Chat.ID
                client.SendMessage(chatID, "Your message has been received!")
            }
        }
        c.Status(200)
    })
    
    client.SetWebhook("https://yourdomain.com/webhook")
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
