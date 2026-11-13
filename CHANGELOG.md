# Changelog

## 1.5.0

- Added automatic retries. A request that fails with a network error, a timeout, or a
  `408`, `429` or `5xx` response is retried up to 2 times with backoff, but only when sending it
  again can't do anything twice. Pass `helo.WithMaxRetries(n)` to `helo.NewHelo` to change how
  many, or `0` to disable them.
- Added a timeout: each attempt is abandoned after 30 seconds (`helo.DefaultTimeout`) and
  retried. Before, the default client had no timeout, so a stalled connection could block
  forever. A client passed to `helo.WithHTTPClient` keeps its own `Timeout`.
- `SendTransactional`, `SendTransactionalBatch`, `SendBroadcast` and `SendBroadcastMessage` now
  generate an idempotency key when `IdempotencyKey` is empty, and send the same key on every
  retry, so a send whose response was lost is never sent twice.

## 1.4.0

- Added sending template management
- Added support for custom tracking domains

## 1.3.1

- Fixed array query parameters whose items are an enum, such as `eventTypes` on `Activity.ListEvents`
  (`helo activity list-events --event-types`): they were sent as `[bounced complained]` instead of
  `bounced,complained`.

## 1.3.0

- Added `NullFields` to `UpdateWebhookRequest`, to send `channelId` or `additionalHeaders` as null, which
  removes a webhook's channel scope or custom headers. Leaving a field unset still leaves it unchanged.

## 1.2.0

- Fixed optional booleans in request types, which could not be set to `false`: with `,omitempty` a
  `false` was dropped from the request, so a webhook couldn't be disabled, nor tracking turned off.
  They are now `*bool`; set them with the new `helo.Bool` helper, e.g. `Enabled: helo.Bool(false)`.
  This changes the type of `Enabled` on `CreateWebhookRequest` and `UpdateWebhookRequest`, `Opens` and
  `Links` on the channel and send request `Tracking` types, and `InlineStyles` on send request templates.

  Fixed `SendMessageResponse`, the type of each result in `SendTransactionalBatch`, which was empty: it
  now carries the fields of both `SendMessageAcceptedResponse` and `SendMessageFailedResponse`.

  `APIError` gained `Status`, `Code`, `RequestID` and `Problem` accessors, and its message now includes
  the API's error code and request ID without repeating the HTTP status.

## 1.1.0

- Added `id` to the `template` of transactional and broadcast send requests, to send a stored template
  instead of an inline `subject`, `html` and `text`.

## 1.0.1

- Release pipeline updates

## 1.0.0

- Initial release
