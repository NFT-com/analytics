# Aggregation API

Aggregation API provides a number of aggregated stats about NFT sales.
Stats describe one of three entities:

- collections
- marketplaces
- NFTs

API often offers two modes for retrieving stats:

- current value for a stat
- historic values for a stat

All historic values accept a `from` and `to` query parameters, representing the date range for which the stat should be returned.
Dates should be provided in the `YYYY-MM-DD` format.

Counter-related stats (like sale or user counts) are typically returned in the following format:

```json
{
    "id": "612ecc22-36ef-4ef7-bb0b-5b864b85d089",
    "value": 123456789
}
```

If the returned stat is currency-related, it is typically represented with a list of currencies.
For instance, trading volume for a collection can be returned as:

```json
{
    "id": "612ecc22-36ef-4ef7-bb0b-5b864b85d089",
    "values": [
        {
            "amount": 1500,
            "currency_id": "e0a9b034-10af-4e5c-ad07-57ee7df67947"
        },
        {
            "amount": 100,
            "currency_id": "00f09a97-5db2-44d1-b666-5e1c872cd0f2"
        },
        {
            "amount": 1000,
            "currency_id": "5c59705d-3aca-4a40-bd27-942830bef64c"
        }
    ]
}
```

Historic data is typically returned in the following format:

```json
[
    {
        "total": 123456789,
        "date": "2022-06-01T00:00:00Z"
    },
    {
        "total": 1234567890,
        "date": "2022-05-31T00:00:00Z"
    }
]
```

## Stats

### Collections Stats

Collections stats include:

- volume
- market cap
- sales
- size
- lowest recent price
- average

#### Collection Volume

Collection volume represents the sum of all trades in a collection.

Current volume for a collection can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/volume'
```

Historic volume for a collection can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/volume/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

Retrieving current collection volume for a number of collections can be done using:

```
curl --location --request POST 'http://localhost:8080/collection/batch/volume' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ids": [
        "<collection-id-1>",
        "<collection-id-2>",
        "<collection-id-3>",
        "<collection-id-4>"
    ]
}'
```

#### Collection Market Cap

Collection market cap represents the sum of all latest prices for NFTs in a collection.
For example, if a collection has two NFTs last sold for 100 USD, market cap for the collection is 200 USD.
If on the following day one of the two NFTs sells for 80 USD, then the market cap for the collection on that day is 180 USD.

Current market cap for a collection can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/market_cap'
```

Historic market cap for a collection can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/market_cap/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

Retrieving current collection market cap for a number of collections can be done using:

```
curl --location --request POST 'http://localhost:8080/collection/batch/market_cap' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ids": [
        "<collection-id-1>",
        "<collection-id-2>",
        "<collection-id-3>",
        "<collection-id-4>"
    ]
}'
```

#### Collection Sales

Collection sales represent the total number of sales involving NFTs in a collection.

Current sale count for a collection can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/sales'
```

Historic sales count for a collection can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/sales/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

#### Collection Size

Collection size represents the total number of NFTs in a collection.

History of the collection size can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/size/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

#### Collection Lowest Price

Collection lowest price represents the lowest price of any NFT in the collection.

Historic lowest prices can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/lowest_price/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

#### Collection Average Price

Average price for a collection represents the average price of an NFT in the collection.

Historic average prices can be retrieved using:

```
curl 'localhost:8080/collection/<collection_id>/size/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

### Marketplace Stats

Marketplace stats include:

- volume
- market cap
- sales
- users

#### Marketplace Volume

Marketplace volume represents the sum of all trades on a marketplace.

Current volume for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/volume'
```

Historic volume for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/volume/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```


#### Marketplace Market Cap

Marketplace market cap represents the sum of all latest prices for NFTs on a marketplace.

Current market cap for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/market_cap'
```

Historic market cap for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/market_cap/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

#### Marketplace Sales

Marketplace sales represent the total number of sales on a marketplace.

Current sale count for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/sales'
```

Historic sales count for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/sales/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

#### Marketplace User Count

Marketplace user count represent the total number of unique users on a marketplace.

Current user count for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/users'
```

Historic user count for a marketplace can be retrieved using:

```
curl 'localhost:8080/marketplace/<marketplace_id>/users/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```

### NFT Stats

NFT stats include:

- price
- average price

#### NFT Price

Current price for an NFT can be retrieved using:

```
curl 'localhost:8080/nft/<nft_id>/price'
```

Historic price for an NFT can be retrieved using:

```
curl 'localhost:8080/nft/<nft_id>/price/history?from=YYYY-MM-DD&to=YYYY-MM-DD'
```
#### NFT Average Price

All time average price for an NFT can be retrieved using:

```
curl 'localhost:8080/nft/<nft_id>/average'
```