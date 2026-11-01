# Helo.Templates

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**List**](Templates.md#list) | **GET** /templates | List all templates |
| [**Create**](Templates.md#create) | **POST** /templates | Create a template |
| [**RenderDraft**](Templates.md#renderdraft) | **POST** /templates/render | Render template content |
| [**Retrieve**](Templates.md#retrieve) | **GET** /templates/{id} | Retrieve a template |
| [**Update**](Templates.md#update) | **PATCH** /templates/{id} | Update a template |
| [**Delete**](Templates.md#delete) | **DELETE** /templates/{id} | Delete a template |
| [**Render**](Templates.md#render) | **POST** /templates/{id}/render | Render a template |


## List

> List(ctx, params) (*PaginatedResponseOfTemplateSummaryResponse, error)

List all templates

Retrieves a paginated list of templates, optionally filtered by name or channel.

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

Create a template

Creates a new template from HTML, MJML, or plain-text content. MJML is compiled to HTML when the template is saved.

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

Render template content

Renders unsaved template content (subject, HTML or MJML, and text) with sample data, without persisting anything.

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

Retrieve a template

Fetches the details and content of a specific template.

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

Update a template

Modifies an existing template. Only the provided fields are changed.

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

Delete a template

Permanently removes a template. This action cannot be undone.

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

Render a template

Renders a saved template with sample data, to preview the output before sending.

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

