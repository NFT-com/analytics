## Goal of this PR

### Description

_High level overview of the changes in this PR..._

### Related issue

_Fixes #_

## Checklist

### Does this PR change GraphQL schema

If this PR changes the GraphQL schema for the Graph API, the following should be updated:

- [ ] [GraphQL Schema documentation](graph/schema/schema.md)
- [ ] GraphQL generated files
- [ ] Postman GraphQL requests/queries
- [ ] GraphQL selectors and arguments in the server code

#### GraphQL selectors and arguments

A number of fields in the GraphQL schema are expensive to compute and only retrieved if requested (e.g. rarity, stats, owners). If any of these fields are touched, the relevant server code should be updated.

### Does this PR change Aggregation API interface

This is the case if this PR changes how the Aggregation APIs serves information.
Adding, renaming or removing endpoints, renaming query parameters or changing their format will impact external users of the API.
This is also the case if the PR changes the data format expected or served by the API - for example,
changing the data model of POST request input data, or the produced JSON response.

GraphQL API optionally consumes data from the Aggregation API, so any changes in the Aggregation API data models should be propagated
to the relevant Graph API code.

All of the following should be updated:
- [ ] Aggregation API [documentation](aggregate/API.md)
- [ ] Postman Aggregation API requests
- [ ] Graph API consumer

### Does this PR change Events API interface

This is the case if this PR changes how the Events APIs serves information.
Adding, renaming or removing endpoints, renaming query parameters or changing their format will impact external users of the API.
This is also the case if the PR changes the data format served by the API.
Postman requests for the Events API and the corresponding tests should be kept up to date.

All of the following should be updated:
- [ ] Events API [documentation](events/API.md)
- [ ] Postman Events API requests
- [ ] Postman Events API tests

### Does this PR change the CLI usage of an executable

This is the case if the PR changes any of the CLI flags required to run the APIs or tool in this repo.

All of the following should be updated:
- [ ] Graph API executable [documentation](cmd/graph-api/README.md)
- [ ] Aggregation API executable [documentation](cmd/aggregation-api/README.md)
- [ ] Events API executable [documentation](cmd/events-api/README.md)
- [ ] Relevant Dockerfile
- [ ] [Docker Compose](docker-compose.yml)

### Does this PR change SQL queries

This is the case if any of the SQL queries for the Graph, Events or the Aggregation API are changed.

- [ ] All data models for SQL results are up to date

### Misc

- [ ] SQL data models are up to date with the [SQL schema](https://github.com/NFT-com/indexer/tree/master/sql)
- [ ] Any deferred tasks have corresponding GitHub issues created
- [ ] PR is against the correct branch
- [ ] PR is labelled appropriately
- [ ] PR is linked to an issue
