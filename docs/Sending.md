# Helo.Sending

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**Transactional**](Sending.md#transactional) | **POST** /send/transactional | Send a transactional email |
| [**TransactionalBatch**](Sending.md#transactionalbatch) | **POST** /send/transactional/batch | Send transactional emails in batch |
| [**Broadcast**](Sending.md#broadcast) | **POST** /send/broadcast | Send a broadcast email |
| [**BroadcastMessage**](Sending.md#broadcastmessage) | **POST** /send/broadcast/message | Send a single broadcast email |


## Transactional

> Transactional(ctx, params) (*SendMessageAcceptedResponse, error)

Send a transactional email

Sends a single transactional email such as receipts, confirmations, or notifications.

### Example

```go Sending_transactional
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.SendMessageRequest{
		Subject: "test-subject",
		Html: "test-html",
		Text: "test-text",
		Tags: []string{"example1", "example2"},
	}
	result, err := client.Sending.Transactional(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## TransactionalBatch

> TransactionalBatch(ctx, params) (*SendMessageBatchResponse, error)

Send transactional emails in batch

Sends multiple transactional emails in a single API request for better performance.

### Example

```go Sending_transactionalBatch
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.SendMessageBatchRequest{
	}
	result, err := client.Sending.TransactionalBatch(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Broadcast

> Broadcast(ctx, params) (*SendBroadcastResponse, error)

Send a broadcast email

Sends a broadcast email to multiple recipients for marketing or announcement purposes.

### Example

```go Sending_broadcast
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.SendBroadcastRequest{
		Tags: []string{"example1", "example2"},
	}
	result, err := client.Sending.Broadcast(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## BroadcastMessage

> BroadcastMessage(ctx, params) (*SendMessageAcceptedResponse, error)

Send a single broadcast email

Sends a single broadcast email message.

### Example

```go Sending_broadcastMessage
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.SendMessageRequest{
		Subject: "test-subject",
		Html: "test-html",
		Text: "test-text",
		Tags: []string{"example1", "example2"},
	}
	result, err := client.Sending.BroadcastMessage(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

