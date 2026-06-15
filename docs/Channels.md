# Helo.Channels

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Channels.md#list) | **GET** /channels | List all channels |
| [**Create**](Channels.md#create) | **POST** /channels | Create a channel |
| [**Retrieve**](Channels.md#retrieve) | **GET** /channels/{id} | Retrieve a channel |
| [**Update**](Channels.md#update) | **PATCH** /channels/{id} | Update a channel |
| [**Delete**](Channels.md#delete) | **DELETE** /channels/{id} | Delete a channel |


## List

> List(ctx, params) (*PaginationResultOfChannelBasicResponse, error)

List all channels

Retrieves a list of all channels accessible to the current user.

### Example

```go Channels_list
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

	params := &helo.ChannelsListParams{
		Limit: 10,
		Offset: 10,
		Name: "example",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
		DeliveryType: "live",
	}
	result, err := client.Channels.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Create

> Create(ctx, params) (*ChannelDetailsResponse, error)

Create a channel

Creates a new communication channel for organizing and routing messages.

### Example

```go Channels_create
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

	params := &helo.CreateChannelRequest{
		Name: "test-name",
		DeliveryType: helo.DeliveryTypeLive,
	}
	result, err := client.Channels.Create(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Retrieve

> Retrieve(ctx, id) (*ChannelDetailsResponse, error)

Retrieve a channel

Fetches the details and configuration of a specific channel.

### Example

```go Channels_retrieve
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

	result, err := client.Channels.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Update

> Update(ctx, id, params) (*ChannelDetailsResponse, error)

Update a channel

Modifies the settings and configuration of an existing channel.

### Example

```go Channels_update
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

	params := &helo.UpdateChannelRequest{
		Name: "test-name",
		DeliveryType: helo.DeliveryTypeLive,
	}
	result, err := client.Channels.Update(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Delete

> Delete(ctx, id) error

Delete a channel

Permanently removes a channel and all associated data.

### Example

```go Channels_delete
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

	if err := client.Channels.Delete(ctx, "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		log.Fatal(err)
	}
}
```

