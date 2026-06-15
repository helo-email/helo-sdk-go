# Helo.Activity

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**ListEvents**](Activity.md#listevents) | **GET** /activity/events | List activity events |
| [**ListMessages**](Activity.md#listmessages) | **GET** /activity/messages | List messages |
| [**RetrieveMessage**](Activity.md#retrievemessage) | **GET** /activity/messages/{id} | Retrieve message details |


## ListEvents

> ListEvents(ctx, params) (*PaginatedEventsResponse, error)

List activity events

Retrieves activity events for messages, including delivery status, opens, clicks, bounces, unsubscribes and complaints.

### Example

```go Activity_listEvents
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	params := &helo.ActivityListEventsParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		MessageID: "550e8400-e29b-41d4-a716-446655440000",
		After: 10,
		StartDate: time.Now(),
		EndDate: time.Now(),
		Limit: 10,
		Recipient: "example",
		Subject: "example",
		Tags: []string{"example1", "example2"},
		MailType: "transactional",
		EventTypes: []helo.EventType{helo.EventTypeAccepted, helo.EventTypeProcessed},
	}
	result, err := client.Activity.ListEvents(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## ListMessages

> ListMessages(ctx, params) (*PaginatedMessagesResponse, error)

List messages

Retrieves a paginated list of sent messages with basic tracking information.

### Example

```go Activity_listMessages
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	params := &helo.ActivityListMessagesParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		After: 10,
		StartDate: time.Now(),
		EndDate: time.Now(),
		Limit: 10,
		Recipient: "example",
		Subject: "example",
		Tags: []string{"example1", "example2"},
		MailType: "transactional",
		Status: "sent",
	}
	result, err := client.Activity.ListMessages(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## RetrieveMessage

> RetrieveMessage(ctx, id) (*MessageDetailsResponse, error)

Retrieve message details

Fetches detailed tracking information for a specific message, including all associated events.

### Example

```go Activity_retrieveMessage
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

	result, err := client.Activity.RetrieveMessage(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

