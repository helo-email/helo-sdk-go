# Helo.WebhookEndpoints

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](WebhookEndpoints.md#list) | **GET** /webhook-endpoints | List all webhook endpoints |
| [**Create**](WebhookEndpoints.md#create) | **POST** /webhook-endpoints | Create a webhook endpoint |
| [**Retrieve**](WebhookEndpoints.md#retrieve) | **GET** /webhook-endpoints/{id} | Retrieve a webhook endpoint |
| [**Update**](WebhookEndpoints.md#update) | **PATCH** /webhook-endpoints/{id} | Update a webhook endpoint |
| [**Delete**](WebhookEndpoints.md#delete) | **DELETE** /webhook-endpoints/{id} | Delete a webhook endpoint |
| [**RegenerateSigningKey**](WebhookEndpoints.md#regeneratesigningkey) | **POST** /webhook-endpoints/{id}/regenerate-signing-key | Regenerate webhook signing key |


## List

> List(ctx, params) (*PaginationResultOfWebhookEndpointResponse, error)

List all webhook endpoints

Retrieves all webhook endpoints configured for the account.

### Example

```go WebhookEndpoints_list
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.WebhookEndpointsListParams{
		Limit: 10,
		Offset: 10,
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.WebhookEndpoints.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Create

> Create(ctx, params) (*WebhookEndpointResponse, error)

Create a webhook endpoint

Registers a new webhook endpoint to receive event notifications.

### Example

```go WebhookEndpoints_create
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.CreateWebhookEndpointRequest{
		URL: "test-url",
		Events: []helo.WebhookEvent{helo.WebhookEventAccepted, helo.WebhookEventProcessed},
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Enabled: true,
	}
	result, err := client.WebhookEndpoints.Create(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Retrieve

> Retrieve(ctx, id) (*WebhookEndpointResponse, error)

Retrieve a webhook endpoint

Fetches the details and configuration of a specific webhook endpoint.

### Example

```go WebhookEndpoints_retrieve
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.WebhookEndpoints.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Update

> Update(ctx, id, params) (*WebhookEndpointResponse, error)

Update a webhook endpoint

Modifies the configuration of an existing webhook endpoint.

### Example

```go WebhookEndpoints_update
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.UpdateWebhookEndpointRequest{
		URL: "test-url",
		Events: []helo.WebhookEvent{helo.WebhookEventAccepted, helo.WebhookEventProcessed},
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Enabled: true,
	}
	result, err := client.WebhookEndpoints.Update(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Delete

> Delete(ctx, id) error

Delete a webhook endpoint

Permanently removes a webhook endpoint.

### Example

```go WebhookEndpoints_delete
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	if err := client.WebhookEndpoints.Delete(ctx, "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		log.Fatal(err)
	}
}
```


## RegenerateSigningKey

> RegenerateSigningKey(ctx, id, ) (*WebhookEndpointResponse, error)

Regenerate webhook signing key

Generates a new signing key for webhook payload verification.

### Example

```go WebhookEndpoints_regenerateSigningKey
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.WebhookEndpoints.RegenerateSigningKey(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

