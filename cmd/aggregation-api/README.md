# Aggregation API

Aggregation API provides a number of aggregated stats about NFT sales.
Stats describe one of three entities:

- collections
- marketplaces
- NFTs

## Usage

```
Usage of aggregation-api:
  -b, --bind string                           bind address for serving requests (default ":8080")
      --enable-query-logging                  enable logging of database queries (default true)
  -e, --events-database string                events database address
      --events-db-connection-limit int        maximum number of connections to events database, -1 for unlimited (default 70)
      --events-db-idle-connection-limit int   maximum number of idle connections to events database (default 20)
  -g, --graph-database string                 graph database address
      --graph-db-connection-limit int         maximum number of connections to graph database, -1 for unlimited (default 70)
      --graph-db-idle-connection-limit int    maximum number of idle connections to graph database (default 20)
  -l, --log-level string                      log level (default "info")
```
