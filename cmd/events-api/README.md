# Events API

Events API serves information about NFT related events occurring on the blockchain.
Events fall into one of two categories: transfers and sales.

Mint and burn events are special instances of transfer events.
For more information about the API schema see [the API documentation](/docs/events-api.md).

## Usage

```
Usage of events-api:
  -s, --batch-size uint        default limit for number of events returned in a single call (default 100)
  -b, --bind string            bind address for serving requests (default ":8080")
  -d, --database string        database address
      --enable-query-logging   enable logging of database queries (default true)
  -l, --log-level string       log level (default "info")
```
