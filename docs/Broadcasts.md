# Helo.Broadcasts

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Broadcasts.md#list) | **GET** /broadcasts | List broadcasts |
| [**Retrieve**](Broadcasts.md#retrieve) | **GET** /broadcasts/{id} | Retrieve a broadcast |
| [**ListFailures**](Broadcasts.md#listfailures) | **GET** /broadcasts/{id}/failures | List broadcast failures |
| [**ListSuppressions**](Broadcasts.md#listsuppressions) | **GET** /broadcasts/{id}/suppressions | List broadcast suppressions |


## List

> List(ctx, params) (*PaginatedResponseOfBroadcast, error)

List broadcasts

Retrieves a paginated list of sent broadcasts with summary statistics.

### Example

```go Broadcasts_list
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.BroadcastsListParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Status: helo.BroadcastStatusAccepted,
		Subject: "example",
		Limit: 10,
		Offset: 10,
	}
	result, err := client.Broadcasts.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Retrieve

> Retrieve(ctx, id) (*BroadcastDetailsResponse, error)

Retrieve a broadcast

Fetches details and statistics for a specific broadcast.

### Example

```go Broadcasts_retrieve
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.Broadcasts.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## ListFailures

> ListFailures(ctx, id) (*PaginatedResponseOfBroadcastFailure, error)

List broadcast failures

Retrieves a list of failed messages for a specific broadcast.

### Example

```go Broadcasts_listFailures
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.Broadcasts.ListFailures(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## ListSuppressions

> ListSuppressions(ctx, id) (*PaginatedResponseOfBroadcastSuppression, error)

List broadcast suppressions

Retrieves a list of suppressed recipients for a specific broadcast.

### Example

```go Broadcasts_listSuppressions
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.Broadcasts.ListSuppressions(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

