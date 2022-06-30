# Getting Started

This guide aims to help users setup the project and run the APIs that are part of this repository.
All docker commands are supposed to be run from the root of the repository.

**Table of Contents**

1. [Requirements](#requirements)
2. [Events API](#events-api)
3. [Aggregation API](#aggregation-api)
4. [Graph API](#graph-api)
5. [Running all APIs with Docker Compose](#running-all-apis-with-docker-compose)
6. [Postman Collections](#postman-collections)

## Requirements

* [Docker](https://docs.docker.com/get-docker/)
* Optional: [Docker Compose](https://docs.docker.com/compose/install/)

## Events API

### Build the Docker Image

To build the docker image for the Events API run:

```console
docker build -t events-api -f ./cmd/events-api/Dockerfile .
```

### Run the Docker Image

To run a docker image inside a container, publishing port `8080` and passing along CLI arguments to the application, we can use `docker run`:

```console
docker run -p 8080:8080 events-api --database "host=host.docker.internal user=db-user password=db-pass dbname=db-name port=5432"
```

### Further Reading

More details about the application's usage and CLI flags can be found at [the Event API's readme file](/cmd/events-api/README.md).

For more details about the API's usage, Postman examples and tests see [the Event API documentation page](/events/API.md).

## Aggregation API

Aggregation API needs to be able to connect to the Events database in order to retrieve sales data.
Aggregation API also needs to be able to connect to the Graph database in order to lookup collection and marketplace chain IDs and contract addresses, as well as token information.

### Build the Docker Image

To build the docker image for the Aggregation API run:

```console
docker build -t aggregation-api -f ./cmd/aggregation-api/Dockerfile .
```

### Running the Docker Image

To run a docker image inside a container, publishing port `8080` and forwarding CLI arguments to the application, we can use `docker run`:

```console
docker run -p 8080:8080 aggregation-api --bind ':8080' --events-database "host=host.docker.internal user=db-user password=db-pass dbname=events-db-name port=5432" --graph-database "host=host.docker.internal user=db-user password=db-pass dbname=graph-db-name port=5432" --events-db-connection-limit 70 --graph-db-connection-limit 70 --log-level debug --enable-query-logging
```

### Further Reading

More details about the application's usage and CLI flags can be found at [the aggregation API's readme file](/cmd/aggregation-api/README.md).

## Graph API

> ⚠️ A running Aggregation API server is a prerequisite for the Graph API.

### Build the Docker Image

To build the docker image for the Graph API run:

```console
docker build -t graph-api -f ./cmd/graph-api/Dockerfile .
```

### Running the Docker Image

To run a docker image inside a container, publishing port `8080` and forwarding CLI arguments to the application, we can use `docker run`:

```console
docker run -p 8080:8080 graph-api --database "host=host.docker.internal user=db-user password=db-pass dbname=db-name port=5432" --enable-playground --bind ':8080'
```

### Further Reading

More details about the application's usage and CLI flags can be found at [the Graph API's readme file](/cmd/graph-api/README.md).

For more details about the API's GraphQL schema, see [/graph/schema/schema.graphql](/graph/schema/schema.graphql).

Generated documentation for the schema can be found at [/graph/schema/schema.md](/graph/schema/schema.md).

## Running all APIs with Docker Compose

It is possible to start all APIs at once with `docker-compose`, using the `docker-compose.yml` file at the root of the repository.
The database connection parameters should instruct the APIs to connect to the database populated by the [indexer](https://github.com/NFT-com/indexer).

### Build the APIs

Building the required images can be done with the following command:

```console
docker-compose build
```

### Starting the APIs

Starting the APIs can be done with the following command:

```console
docker-compose up -d
```

> 🚧 Since the Aggregation API is not yet implemented, the corresponding service will exit immediately.

## Postman Collections

A number of Postman collections testing the APIs, as well as instructions on how to run them, can be found in the [/postman](/postman) directory.