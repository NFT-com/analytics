# Graph API

## Usage

```
Usage of graph-api:
  -b, --bind string              bind address for serving requests (default ":8080")
  -d, --database string          database address
  -l, --log-level string         log level (default "info")
  -p, --playground-path string   path for GraphQL playground (default "/")
      --query-complexity int     GraphQL query complexity limit
      --enable-playground        enable GraphQL playground
      --enable-query-logging     enable logging of database queries (default true)
```

## Database Address - Data Source Name

Data Source Name (DSN) is the string specified describing how the connection to the database should be established.
Format of the string is the following:

```
host=localhost user=database-user password=password dbname=database-name port=5432 sslmode=disable
```

## GraphQL Playground

The API provides a command line option to enable [GraphQL Playground](https://github.com/graphql/graphql-playground), and IDE for GraphQL development.
By default, GraphQL Playground is disabled, but can be enabled using the `--enable-playground` command line option.
The URL where GraphQL Playground is hosted can be set using the `--playground-path` command line option.

## GraphQL Query Complexity

The API provides a command line option to limit [query complexity](https://gqlgen.com/reference/complexity/).
Default value is 0, meaning that the query complexity is not limited.
For a non-zero limit, each requested field and each level of depth adds one to the overall query complexity.
When the limit is reached, the query is rejected and an error is returned.
