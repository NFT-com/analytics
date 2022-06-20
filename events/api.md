# Events API

Events API is a REST API serving data about NFT-related events occurring on a blockchain.
Events are served in a reverse chronological order, meaning most recent events are returned first.

## Endpoints

In order to have greater flexibility since different event types can require different filters, event listing is done using multiple endpoints.
Supported endpoints are:

- `/transfers/`
- `/sales/`

Listing all events for a single NFT is done by issuing requests to individual endpoints and merging the results.

### Transfers

Listing transfer events is done by issuing a `GET` request to the `/transfers` endpoint.
The following filters are available for transfer events:

- `chain_id` — chain ID
- `collection_address` — address of the collection
- `token_id` — ID of the token, as seen on chain
- `transaction_hash` — transaction hash
- `start_timestamp` — start timestamp in RFC3339 format (inclusive)
- `end_timestamp` — end timestamp in RFC3339 format (inclusive)
- `start_height` — start height for block range (inclusive)
- `end_height` — end height for block range (inclusive)
- `sender_address` — address of the token sender
- `receiver_address` — address of the token receiver 

### Sales

Listing sale events is done by issuing a `GET` request to the `/sales/` endpoint.
The following filters are available for sale events:

- `chain_id` — chain ID
- `marketplace_address` — address of the marketplace
- `collection_address` — address of the collection
- `token_id` - ID of the token, as seen on chain
- `transaction_hash` — transaction hash
- `start_timestamp` — start timestamp in RFC3339 format (inclusive)
- `end_timestamp` — end timestamp in RFC3339 format (inclusive)
- `start_height` — start height for block range (inclusive)
- `end_height` — end height for block range (inclusive)
- `seller_address` — address of the token seller
- `buyer_address` — address of the token buyer
- `start_price` — start price for price range (inclusive)
- `end_price` — end price for price range (inclusive)

## Pagination

Since the number of events fitting a search criteria can be large, events are returned in batches.
Default number of events returned in a single batch is configured on the back-end side via a command line flag.

When a given search criteria has more events than what is returned in a single batch, a `next_page` token is included in the API response:

```json
{
  "events": [],
  "next_page": "MTAwMDAwMDk5MDow"
}
```

This token should be provided in the subsequent request as a `page` query parameter, along with the original search parameters — e.g. `curl -X GET localhost:8080/transfers/?collection=collectionID&token_id=tokenID&page=MTAwMDAwMDk5MDow`.
The process should be repeated until there are no more pages to process, at which point the `next_page` token is omitted from the response.

## Examples and Tests

Postman collections with query examples can be found at [/resources/postman/](/resources/postman/).

Each collection and each request have a number of tests associated with them, which can be executed by using the [Postman Collection Runner](https://learning.postman.com/docs/running-collections/intro-to-collection-runs/).
Tests verify that all requests succeeded with a `200 OK` status code, and that the responses received are the expected ones.
