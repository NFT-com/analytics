# Events API

Events API is a REST API serving time series data relating to NFT-related events occurring on a blockchain.

## Endpoints

In order to have greater flexibility since different event types can require different filters, event listing is done using multiple endpoints.
Supported endpoints are:

- `/mints/`
- `/transfers/`
- `/burns/`
- `/sales/`

Listing all events for a single NFT is done by issuing requests to individual endpoints and merging the results.

### Filters

Some filters will apply to all endpoints.
It should be possible to retrieve all events by specifying:

- chain
- collection
- marketplace
- NFT (by token ID)
- time range (in the form of `start` and `end` times)

Some filters are specific for a single endpoint, for example, the `/transfers/` endpoint could have `from` and `to` parameters (Ethereum addresses) as well as the price, while the `/burns/` endpoint could have the `owner` parameter.

### Paging

Number of events for a specific query can be rather large.
One frequent paging mechanism is *cursor based paging* where the token for the next page is returned in the response payload, and that token is used to retrieve the next page of results.

### Query examples

List all `mint` events on Ethereum in a date range:

```
curl events.nft.com/mints/?chain=ethereum&start=1234560&end=1234567
```

List all `transfer` events on Ethereum between two specific addresses:

```
curl events.nft.com/transfers/?start=1234560&end=1234567&from=<addressA>&to=<addressB>
```

List all `transfer` Ethereum events for Opensea marketplace relating to the CryptoPunks collection:

```
curl events.nft.com/transfers/?chain=ethereum&collection=cryptopunks&marketplace=opensea
```