# Events Postman Collections

## Tests

For each request in a collection, a check is made whether the request succeeded with a `200 OK` HTTP status, as well as if the request returned at least one result (transfer or sale event).

Besides this, each request has a number of other checks associated with it.
For instance, if the HTTP request requests all events in a block range, it is checked that the `block_number` of returned events is within the specified block range.

Requests have a number of parameters that should be changed depending on the underlying data served by the API - e.g. block range, indexed collections etc.
