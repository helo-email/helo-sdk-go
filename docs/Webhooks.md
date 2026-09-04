# Helo.Webhooks

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Webhooks.md#list) | **GET** /webhooks | List all webhooks |
| [**Create**](Webhooks.md#create) | **POST** /webhooks | Create a webhook |
| [**Retrieve**](Webhooks.md#retrieve) | **GET** /webhooks/{id} | Retrieve a webhook |
| [**Update**](Webhooks.md#update) | **PATCH** /webhooks/{id} | Update a webhook |
| [**Delete**](Webhooks.md#delete) | **DELETE** /webhooks/{id} | Delete a webhook |
| [**RegenerateSigningKey**](Webhooks.md#regeneratesigningkey) | **POST** /webhooks/{id}/regenerate-signing-key | Regenerate webhook signing key |


## List

> List(ctx, params) (*PaginationResultOfWebhookResponse, error)

List all webhooks

Retrieves all webhooks configured for the account.

### Example

```go Webhooks_list
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

	params := &helo.WebhooksListParams{
		Limit: 10,
		Offset: 10,
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Webhooks.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Create

> Create(ctx, params) (*WebhookResponse, error)

Create a webhook

Registers a new webhook to receive event notifications.

### Example

```go Webhooks_create
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

	params := &helo.CreateWebhookRequest{
		URL: "test-url",
		Events: []helo.WebhookEvent{helo.WebhookEventMessageAccepted, helo.WebhookEventMessageProcessed},
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Enabled: true,
	}
	result, err := client.Webhooks.Create(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Retrieve

> Retrieve(ctx, id) (*WebhookResponse, error)

Retrieve a webhook

Fetches the details and configuration of a specific webhook.

### Example

```go Webhooks_retrieve
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

	result, err := client.Webhooks.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Update

> Update(ctx, id, params) (*WebhookResponse, error)

Update a webhook

Modifies an existing webhook by ID.

### Example

```go Webhooks_update
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

	params := &helo.UpdateWebhookRequest{
		URL: "test-url",
		Events: []helo.WebhookEvent{helo.WebhookEventMessageAccepted, helo.WebhookEventMessageProcessed},
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Enabled: true,
	}
	result, err := client.Webhooks.Update(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Delete

> Delete(ctx, id) error

Delete a webhook

Permanently removes a webhook.

### Example

```go Webhooks_delete
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

	if err := client.Webhooks.Delete(ctx, "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		log.Fatal(err)
	}
}
```


## RegenerateSigningKey

> RegenerateSigningKey(ctx, id) (*WebhookResponse, error)

Regenerate webhook signing key

Regenerate the signing key used for the webhook signature. This operation replaces the old key.

### Example

```go Webhooks_regenerateSigningKey
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

	result, err := client.Webhooks.RegenerateSigningKey(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

