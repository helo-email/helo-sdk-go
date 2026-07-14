# Helo.Broadcasts

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Broadcasts.md#list) | **GET** /broadcasts | List broadcasts |
| [**Retrieve**](Broadcasts.md#retrieve) | **GET** /broadcasts/{id} | Retrieve a broadcast |
| [**ListFailures**](Broadcasts.md#listfailures) | **GET** /broadcasts/{id}/failures | List failed broadcast messages |
| [**ListSuppressions**](Broadcasts.md#listsuppressions) | **GET** /broadcasts/{id}/suppressions | List broadcast suppressed recipients |


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
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
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
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))
	ctx := context.Background()

	result, err := client.Broadcasts.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## ListFailures

> ListFailures(ctx, id, params) (*PaginatedResponseOfBroadcastFailure, error)

List failed broadcast messages

Returns messages that could not be delivered due to permanent errors (e.g. invalid addresses, domain issues). Transient errors that were retried successfully do not appear here.

### Example

```go Broadcasts_listFailures
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

	params := &helo.BroadcastsListFailuresParams{
		Limit: 10,
		Offset: 10,
	}
	result, err := client.Broadcasts.ListFailures(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## ListSuppressions

> ListSuppressions(ctx, id, params) (*PaginatedResponseOfBroadcastSuppression, error)

List broadcast suppressed recipients

Returns recipients that were skipped because they appear on a suppression list (e.g. previous bounces or unsubscribes).

### Example

```go Broadcasts_listSuppressions
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

	params := &helo.BroadcastsListSuppressionsParams{
		Limit: 10,
		Offset: 10,
	}
	result, err := client.Broadcasts.ListSuppressions(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

