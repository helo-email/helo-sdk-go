# Helo.Statistics

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**RetrieveHourly**](Statistics.md#retrievehourly) | **GET** /activity/statistics/hourly | Retrieve hourly statistics |
| [**RetrieveDaily**](Statistics.md#retrievedaily) | **GET** /activity/statistics/daily | Retrieve daily statistics |
| [**RetrieveTotals**](Statistics.md#retrievetotals) | **GET** /activity/statistics/totals | Retrieve all time statistics |


## RetrieveHourly

> RetrieveHourly(ctx, params) (*StatisticsHourlyResponse, error)

Retrieve hourly statistics

Fetches hourly aggregated statistics.

### Example

```go Statistics_retrieveHourly
package main

import (
	"context"
	"log"
	"time"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_ACCESS_TOKEN")
	ctx := context.Background()

	params := &helo.StatisticsRetrieveHourlyParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		From: time.Now(),
		To: time.Now(),
		Tags: []string{"example1", "example2"},
	}
	result, err := client.Statistics.RetrieveHourly(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## RetrieveDaily

> RetrieveDaily(ctx, params) (*StatisticsDailyResponse, error)

Retrieve daily statistics

Fetches daily aggregated statistics.

### Example

```go Statistics_retrieveDaily
package main

import (
	"context"
	"log"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_ACCESS_TOKEN")
	ctx := context.Background()

	params := &helo.StatisticsRetrieveDailyParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		From: "2024-01-01",
		To: "2024-01-01",
		Tags: []string{"example1", "example2"},
		Timezone: "America/New_York",
	}
	result, err := client.Statistics.RetrieveDaily(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```


## RetrieveTotals

> RetrieveTotals(ctx, params) (*StatisticsTotalsResponse, error)

Retrieve all time statistics

Fetches cumulative statistics.

### Example

```go Statistics_retrieveTotals
package main

import (
	"context"
	"log"
	"time"
	"github.com/helo-email/helo-sdk-go"
)

func main() {
	client := helo.NewHelo("YOUR_ACCESS_TOKEN")
	ctx := context.Background()

	params := &helo.StatisticsRetrieveTotalsParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		From: time.Now(),
		To: time.Now(),
		Tags: []string{"example1", "example2"},
	}
	result, err := client.Statistics.RetrieveTotals(ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	_ = result
}
```

