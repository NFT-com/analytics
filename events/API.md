# Events API

Events API is a REST API serving time series data relating to NFT-related events occurring on a blockchain.
Events are served in a reverse chronological order, meaning most recent events are returned first.

## Endpoints

In order to have greater flexibility since different event types can require different filters, event listing is done using multiple endpoints.
Supported endpoints are:

- `/mints/`
- `/transfers/`
- `/burns/`
- `/sales/`

Listing all events for a single NFT is done by issuing requests to individual endpoints and merging the results.

### Mints

Listing mint events is done by issuing a `GET` request to the `/mints/` endpoint.
The following filters are available for mint events:

- `collection` - UUID of the collection
- `token_id` - Non-Fungible Token ID, as seen on the chain
- `transaction` - transaction ID
- `start` - start time in RFC3339 format (inclusive)
- `end` - end time in RFC3339 format (inclusive)
- `block_start` - string representation of the block number in decimal format (inclusive)
- `block_end` - string representation of the block number in decimal format (inclusive)
- `owner` - address of the token owner

### Transfers

Listing transfer events is done by issuing a `GET` request to the `/transfers` endpoint.
The following filters are available for transfer events:

- `collection` - UUID of the collection
- `token_id` - Non-Fungible Token ID, as seen on the chain
-- `transaction` - transaction ID
- `start` - start time in RFC3339 format (inclusive)
- `end` - end time in RFC3339 format (inclusive)
- `block_start` - string representation of the block number in decimal format (inclusive)
- `block_end` - string representation of the block number in decimal format (inclusive)
- `from` - address of the token sender
- `to` - address of the token receiver 

### Burns

Listing burn events is done by issuing a `GET` request to the `/burns/` endpoint.
The following filters are available for burn events:

- `collection` - UUID of the collection
- `token_id` - Non-Fungible Token ID, as seen on the chain
- `transaction` - transaction ID
- `start` - start time in RFC3339 format (inclusive)
- `end` - end time in RFC3339 format (inclusive)
- `block_start` - string representation of the block number in decimal format (inclusive)
- `block_end` - string representation of the block number in decimal format (inclusive)

### Sales

Listing sale events is done by issuing a `GET` request to the `/sales/` endpoint.
The following filters are available for sale events:

- `marketplace` - UUID of the marketplace
- `transaction` - transaction ID
- `start` - start time in RFC3339 format (inclusive)
- `end` - end time in RFC3339 format (inclusive)
- `block_start` - string representation of the block number in decimal format (inclusive)
- `block_end` - string representation of the block number in decimal format (inclusive)
- `seller` - address of the token seller
- `buyer` - address of the token buyer
- `price` - price for the token

## Pagination

Since the number of events fitting a search criteria can be large, events are returned in batches.
Default number of events returned in a single batch is configured on the back-end side via a CLI flag.

When a given search criteria has more events than what is returned in a single batch, a `next_page` token is included in the API response:

```json
{
  "events": [],
  "next_page": "MTAwMDAwMDk5MDow"
}
```

This token should be provided in the subsequent request as a `page` query parameter, along with the original search parameters - e.g. `curl -X GET localhost:8080/transfers/?collection=collectionID&token_id=tokenID&page=MTAwMDAwMDk5MDow`.
The process should be repeated until there are no more pages to process, at which point the `next_page` token will be omitted from the response.

## Examples and Tests

Postman collections with query examples can be found at [/resources/postman/](/resources/postman/).

Each collection and each request have a number of tests associated with them, which can be executed by using the [Postman Collection Runner](https://learning.postman.com/docs/running-collections/intro-to-collection-runs/).
Tests verify that all requests succeeded with a `200 OK` status code, as well as that the output has the expected output.
