# Graph API Pagination System

This document describes the Graph API pagination system.

The pagination system uses two arguments: `first` and `after`.

The `first` argument is a positive integer that should correspond to the number of elements to be returned as part of the page.
The `after` argument is optional and should be set to retrieve subsequent pages, with the identifier of the last retrieved element, in order to get the _first X_ elements _after_ the last one that was retrieved.

## Examples

### Retrieving the First `100` NFTs

```graphql
query ($id: ID!) {
  collection(id: $id) {
    id
    name
    nfts(first: 100) {
      edges {
        node {
          id
        }
        cursor
      }
      pageInfo {
        hasNextPage
      }
    }
  }
}
```

This call should result in a list of NFTs like the following:

```json
{
  "data": {
    "collection": {
      "id": "17d5f376-954f-4167-ac7c-0007df5efa62",
      "name": "",
      "description": "Cometh is a DeFi powered game with yield generating NFT. Get spaceships, explore the galaxy and earn tokens.",
      "nfts": {
        "edges": [
          {
            "node": {
              "id": "001a75da-6112-1aef-79a2-b1ff4f3c04f0",
              "token_id": "1000649",
              "uri": "https://nft.service.cometh.io/1000649"
            },
            "cursor": "MDAxYTc1ZGEtNjExMi0xYWVmLTc5YTItYjFmZjRmM2MwNGYw"
          },
          {
            "node": {
              "id": "00277453-d70d-a8ae-b23d-49dcec4c8516",
              "token_id": "7000374",
              "uri": "https://nft.service.cometh.io/7000374"
            },
            "cursor": "MDAyNzc0NTMtZDcwZC1hOGFlLWIyM2QtNDlkY2VjNGM4NTE2"
          }
        ],
        "pageInfo": {
          "hasNextPage": true
        }
      }
    }
  }
}
```

If the retrieved data has `data.collection.nfts.pageInfo.hasNextPage` set to true, then  take note of the `cursor` value for the last element of the `data.collection.nfts.edges` array.
This value is what needs to be set as the `after` argument in the next call in order to retrieve the next page.

### Retrieving the Next Page

Following the steps described in [Retrieving the First `100` NFTs](#retrieving-the-first-100-nfts), simply pass the value of `cursor` as the `after` argument like such:

```graphql
query ($id: ID!) {
  collection(id: $id) {
    id
    name
    nfts(first: 100, after: "MDAyNzc0NTMtZDcwZC1hOGFlLWIyM2QtNDlkY2VjNGM4NTE2") {
      edges {
        node {
          id
        }
        cursor
      }
      pageInfo {
        hasNextPage
      }
    }
  }
}
```
