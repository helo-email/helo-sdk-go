# Helo.Webhooks

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**ListForChannel**](Webhooks.md#listforchannel) | **GET** /app/channels/{id}/webhooks | ListForChannel operation |
| [**List**](Webhooks.md#list) | **GET** /webhooks | List operation |
| [**Create**](Webhooks.md#create) | **POST** /webhooks | Create operation |
| [**Retrieve**](Webhooks.md#retrieve) | **GET** /webhooks/{id} | Retrieve operation |
| [**Update**](Webhooks.md#update) | **PATCH** /webhooks/{id} | Update operation |
| [**Delete**](Webhooks.md#delete) | **DELETE** /webhooks/{id} | Delete operation |
| [**RegenerateSigningKey**](Webhooks.md#regeneratesigningkey) | **POST** /webhooks/{id}/regenerate-signing-key | RegenerateSigningKey operation |


## ListForChannel

> ListForChannel(ctx, id) (*WebhooksResponse, error)

ListForChannel operation

Retrieve all webhooks applicable to a specific channel.

### Example

```go Webhooks_listForChannel
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

	result, err := client.Webhooks.ListForChannel(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## List

> List(ctx, params) (*PaginationResultOfWebhookResponse, error)

List operation

List webhooks.

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

Create operation

Create a new webhook.

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
		Events: []helo.WebhookEvent{helo.WebhookEventAccepted, helo.WebhookEventProcessed},
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

Retrieve operation

Retrieve a single webhook by ID.

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

Update operation

Update an existing webhook.

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
		Events: []helo.WebhookEvent{helo.WebhookEventAccepted, helo.WebhookEventProcessed},
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

Delete operation

Delete a webhook by ID.

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

RegenerateSigningKey operation

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

