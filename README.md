# Analytics

The analytics project contains binaries that expose three APIs that use databases built by the [NFT.com indexer](https://github.com/NFT-com/indexer).

The events API requires access to the Events database, the graph API requires access to the Graph database and the aggregation API requires both.

## Getting Started

Follow the [getting started guide](./docs/getting-started.md) to learn the requirements of the binaries, how to build and run them, and how to test the APIs using postman collections.

## API Documentation

Here are the documents that describe the APIs within this project.

* [Events API Documentation](./docs/events-api.md)
* [Graph API Documentation](./graph/schema/schema.md)
  * [Pagination System](./docs/graph-api-pagination.md)
* [Aggregation API Documentation](./docs/aggregation-api.md)