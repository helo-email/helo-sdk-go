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

## Webhook signature verification

Verify incoming webhooks against the raw (unread) request body:

```go
body, err := io.ReadAll(r.Body)
if err != nil {
	http.Error(w, "bad request", http.StatusBadRequest)
	return
}

if err := helo.VerifyWebhookSignature(r.Header.Get("X-Helo-Webhook-Signature"), body, signingKey); err != nil {
	http.Error(w, "invalid signature", http.StatusBadRequest)
	return
}
```

The header may carry more than one signature (`t=...,v1=...,v2=...`) while a new signing
scheme is being rolled out. Verification uses the newest version present that this SDK
supports — see `helo.SupportedWebhookSignatureVersions` — and ignores the rest,
so a rollout does not break receivers that have not upgraded yet.

`VerifyWebhookSignature` returns a sentinel error describing the rejection, so you can treat a
clock-skew retry differently from a genuinely bad signature:

```go
switch {
case errors.Is(err, helo.ErrSignatureTimestampSkew):
	// delivery is stale — safe to retry, and worth checking clock drift
case errors.Is(err, helo.ErrSignatureMismatch):
	// wrong signing key, or the body was modified in transit
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
- [Webhooks](docs/Webhooks.md)
