# Helo Go SDK

Helo API

## Installation

```bash
go get github.com/helo-email/helo-sdk-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo(os.Getenv("HELO_API_KEY"))

	ctx := context.Background()
	result, err := client.Channels.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
```

## Resources

- [Channels](docs/Channels.md)
- [Activity](docs/Activity.md)
- [Domains](docs/Domains.md)
- [Sending](docs/Sending.md)
- [Broadcasts](docs/Broadcasts.md)
- [Statistics](docs/Statistics.md)
- [Suppressions](docs/Suppressions.md)
- [WebhookEndpoints](docs/WebhookEndpoints.md)
