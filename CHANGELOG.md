# Changelog

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
