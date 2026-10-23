# Helo.Templates

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Templates.md#list) | **GET** /templates | List operation |
| [**Create**](Templates.md#create) | **POST** /templates | Create operation |
| [**RenderDraft**](Templates.md#renderdraft) | **POST** /templates/render | RenderDraft operation |
| [**Retrieve**](Templates.md#retrieve) | **GET** /templates/{id} | Retrieve operation |
| [**Update**](Templates.md#update) | **PATCH** /templates/{id} | Update operation |
| [**Delete**](Templates.md#delete) | **DELETE** /templates/{id} | Delete operation |
| [**Render**](Templates.md#render) | **POST** /templates/{id}/render | Render operation |


## List

> List(ctx, params) (*PaginatedResponseOfTemplateSummaryResponse, error)

List operation

Fetch a paginated list of templates.

### Example

```go Templates_list
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

	params := &helo.TemplatesListParams{
		Limit: 10,
		Offset: 10,
		Name: "example",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Templates.List(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Create

> Create(ctx, params) (*TemplateResponse, error)

Create operation

Create a new template.

### Example

```go Templates_create
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

	params := &helo.CreateTemplateRequest{
		Name: "test-name",
		Subject: "test-subject",
		Html: "test-html",
		Mjml: "test-mjml",
		Text: "test-text",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Templates.Create(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## RenderDraft

> RenderDraft(ctx, params) (*RenderTemplateResponse, error)

RenderDraft operation

Render ad-hoc template content (Subject/Html/Text) with sample data, without persisting anything.

### Example

```go Templates_renderDraft
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

	params := &helo.RenderDraftRequest{
		Subject: "test-subject",
		Html: "test-html",
		Mjml: "test-mjml",
		Text: "test-text",
		InlineStyles: helo.Bool(true),
	}
	result, err := client.Templates.RenderDraft(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Retrieve

> Retrieve(ctx, id) (*TemplateResponse, error)

Retrieve operation

Fetch a single template.

### Example

```go Templates_retrieve
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

	result, err := client.Templates.Retrieve(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Update

> Update(ctx, id, params) (*TemplateResponse, error)

Update operation

Update an existing template.

### Example

```go Templates_update
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

	params := &helo.UpdateTemplateRequest{
		Name: "test-name",
		Subject: "test-subject",
		Html: "test-html",
		Mjml: "test-mjml",
		Text: "test-text",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
		Version: 123,
	}
	result, err := client.Templates.Update(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## Delete

> Delete(ctx, id) error

Delete operation

Delete a template.

### Example

```go Templates_delete
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

	if err := client.Templates.Delete(ctx, "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		log.Fatal(err)
	}
}
```


## Render

> Render(ctx, id, params) (*RenderTemplateResponse, error)

Render operation

Render a saved template with sample data.

### Example

```go Templates_render
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

	params := &helo.RenderTemplateRequest{
		InlineStyles: helo.Bool(true),
	}
	result, err := client.Templates.Render(ctx, "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

