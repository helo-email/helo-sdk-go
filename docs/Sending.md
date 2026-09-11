# Helo.Sending

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**SendTransactional**](Sending.md#sendtransactional) | **POST** /send/transactional | Send a transactional email |
| [**SendTransactionalBatch**](Sending.md#sendtransactionalbatch) | **POST** /send/transactional/batch | Send transactional emails in batch |
| [**SendBroadcast**](Sending.md#sendbroadcast) | **POST** /send/broadcast | SendBroadcast operation |
| [**SendBroadcastMessage**](Sending.md#sendbroadcastmessage) | **POST** /send/broadcast/message | Send a single broadcast email |


## SendTransactional

> SendTransactional(ctx, params, opts) (*SendMessageAcceptedResponse, error)

Send a transactional email

Sends a single transactional email such as receipts, confirmations, or notifications.

### Example

```go Sending_sendTransactional
package main

import (
	"context"
	"log"
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	params := &helo.SendMessageRequest{
		From: helo.MailAddress{Email: "from@yourdomain.com", Name: "From name"},
		To: []helo.MailAddress{{Email: "to@example.com", Name: "To name"}},
		Subject: "Hello from Helo",
		Html: "<html><body><h1>Hi there, new friend.</h1><p>This is a test message, delivered with <3 by Helo. </p></body></html>",
		Text: "This is a test message, delivered with <3 by Helo.",
		Tags: []string{"welcome", "onboarding"},
	}
	opts := &helo.SendingSendTransactionalOptions{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendTransactional(ctx, params, opts)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## SendTransactionalBatch

> SendTransactionalBatch(ctx, params, opts) (*SendMessageBatchResponse, error)

Send transactional emails in batch

Sends multiple transactional emails in a single API request for better performance.

### Example

```go Sending_sendTransactionalBatch
package main

import (
	"context"
	"log"
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	params := &helo.SendMessageBatchRequest{
		Requests: []helo.SendMessageRequest{{From: helo.MailAddress{Email: "from@yourdomain.com", Name: "From name"}, To: []helo.MailAddress{{Email: "to@example.com", Name: "To name"}}, Subject: "Hello from Helo", Html: "<html><body><h1>Hi there, new friend.</h1><p>This is a test message, delivered with <3 by Helo. </p></body></html>", Text: "This is a test message, delivered with <3 by Helo.", Tags: []string{"welcome", "onboarding"}}},
	}
	opts := &helo.SendingSendTransactionalBatchOptions{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendTransactionalBatch(ctx, params, opts)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## SendBroadcast

> SendBroadcast(ctx, params, opts) (*SendBroadcastResponse, error)

SendBroadcast operation

### Example

```go Sending_sendBroadcast
package main

import (
	"context"
	"log"
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	params := &helo.SendBroadcastRequest{
		From: helo.MailAddress{Email: "test@example.com", Name: "test-name"},
		Template: helo.SendBroadcastRequestTemplate{Subject: "test-subject", Html: "test-html", Text: "test-text", InlineStyles: true},
		Tags: []string{"example1", "example2"},
		Messages: []helo.SendBroadcastRequestMessage{{To: []helo.MailAddress{{Email: "test@example.com", Name: "test-name"}}, Tags: []string{"example1", "example2"}}},
	}
	opts := &helo.SendingSendBroadcastOptions{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendBroadcast(ctx, params, opts)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## SendBroadcastMessage

> SendBroadcastMessage(ctx, params, opts) (*SendMessageAcceptedResponse, error)

Send a single broadcast email

Sends a single broadcast email message.

### Example

```go Sending_sendBroadcastMessage
package main

import (
	"context"
	"log"
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	params := &helo.SendMessageRequest{
		From: helo.MailAddress{Email: "from@yourdomain.com", Name: "From name"},
		To: []helo.MailAddress{{Email: "to@example.com", Name: "To name"}},
		Subject: "Hello from Helo",
		Html: "<html><body><h1>Hi there, new friend.</h1><p>This is a test message, delivered with <3 by Helo. </p></body></html>",
		Text: "This is a test message, delivered with <3 by Helo.",
		Tags: []string{"welcome", "onboarding"},
	}
	opts := &helo.SendingSendBroadcastMessageOptions{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendBroadcastMessage(ctx, params, opts)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

