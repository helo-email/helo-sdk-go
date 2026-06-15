# Helo.Suppressions

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Suppressions.md#list) | **GET** /suppressions | List suppressions |
| [**Create**](Suppressions.md#create) | **POST** /suppressions | Create suppressions |
| [**Remove**](Suppressions.md#remove) | **POST** /suppressions/remove | Remove suppressions |


## List

> List(ctx, params) (*PaginatedResponseOfSuppressionResponse, error)

List suppressions

Retrieves a list of suppressed email addresses for a channel.

### Example

```go Suppressions_list
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

	params := &helo.SuppressionsListParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		MailType: helo.MailTypeTransactional,
		Reason: helo.SuppressionReasonBounce,
		Email: "test@example.com",
		Limit: 10,
		Offset: 10,
	}
	result, err := client.Suppressions.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Create

> Create(ctx, params) (*CreateSuppressionsResponse, error)

Create suppressions

Adds email addresses to the suppression list to prevent future sends.

### Example

```go Suppressions_create
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

	params := &helo.CreateSuppressionsRequest{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		MailType: helo.MailTypeTransactional,
		Emails: []string{"example1", "example2"},
	}
	result, err := client.Suppressions.Create(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Remove

> Remove(ctx, params) (*RemoveSuppressionsResponse, error)

Remove suppressions

Removes email addresses from the suppression list to allow future sends.

### Example

```go Suppressions_remove
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

	params := &helo.RemoveSuppressionsRequest{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		MailType: helo.MailTypeTransactional,
		Emails: []string{"example1", "example2"},
	}
	result, err := client.Suppressions.Remove(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

