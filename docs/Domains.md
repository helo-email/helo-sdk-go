# Helo.Domains

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Domains.md#list) | **GET** /domains | List all domains |
| [**Create**](Domains.md#create) | **POST** /domains | Create a domain |
| [**Retrieve**](Domains.md#retrieve) | **GET** /domains/{id} | Retrieve a domain |
| [**Update**](Domains.md#update) | **PATCH** /domains/{id} | Update a domain |
| [**Delete**](Domains.md#delete) | **DELETE** /domains/{id} | Delete a domain |
| [**Verify**](Domains.md#verify) | **POST** /domains/{id}/verify | Verify a domain |
| [**RotateKey**](Domains.md#rotatekey) | **POST** /domains/{id}/rotate-key | Rotate a domain key |


## List

> List(ctx, params) (*PaginatedResponseOfDomainResponse, error)

List all domains

Retrieves all domains associated with the current account, including their verification status.

### Example

```go Domains_list
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.DomainsListParams{
		Limit: 10,
		Offset: 10,
		Name: "example",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Domains.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Create

> Create(ctx, params) (*DomainWithDnsResponse, error)

Create a domain

Registers a new domain for sending emails. The domain must be verified before it can be used.

### Example

```go Domains_create
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.CreateDomainRequest{
		Name: "test-name",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Domains.Create(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Retrieve

> Retrieve(ctx, id) (*DomainWithDnsResponse, error)

Retrieve a domain

Gets detailed information about a specific domain, including verification status and configuration.

### Example

```go Domains_retrieve
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.Domains.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Update

> Update(ctx, id, params) (*DomainResponse, error)

Update a domain

Modifies the configuration settings of an existing domain.

### Example

```go Domains_update
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	params := &helo.UpdateDomainRequest{
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Domains.Update(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Delete

> Delete(ctx, id) error

Delete a domain

Removes a domain from the account. This will stop all email sending from this domain.

### Example

```go Domains_delete
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	if err := client.Domains.Delete(ctx, "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		log.Fatal(err)
	}
}
```


## Verify

> Verify(ctx, id) (*DnsRecordsResponse, error)

Verify a domain

Initiates the domain verification process by checking DNS records.

### Example

```go Domains_verify
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.Domains.Verify(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## RotateKey

> RotateKey(ctx, id) (*DnsRecordResponse, error)

Rotate a domain key

Generates new DKIM keys for the domain. This is recommended for security best practices.

### Example

```go Domains_rotateKey
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_API_KEY")
	ctx := context.Background()

	result, err := client.Domains.RotateKey(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

