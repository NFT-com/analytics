# Postman collections

The Postman collection file provides requests for each of the Analytic APIs: Graph, Events and Aggregation.
These sample requests help showcase the different use case behind each API.
While some input parameters are specified directly in the request, a handful of other variables are configurable on the Collection level. 

## Running tests from CLI

It is possible to run the collections and associated tests from the command line using a tool like [Newman](https://learning.postman.com/docs/running-collections/using-newman-cli/command-line-integration-with-newman/).

### Installation

Newman CLI tool can be installed with the following command:

```console
$ npm install -g newman
```

### Examples

#### Running The Entire Collection

All requests in a collection can be ran using the following command.

```console
newman run <path-to-postman-collection>
```

#### Running a Specific Folder/request

Running a subset of Postman requests can be done by specifying the `--folder` parameter.
Unfortunately support for nested folders/requests is limited, but if the folder/request name is unique within the collection, it is possible to run a specific request.
For example, to run the `Get Collection` request in the `Graph` folder, specifying `--folder 'Get Collection' will suffice.

```console
newman run <path-to-postman-collection> --folder 'Get Collection'
```

#### Specifying variables

Newman will use the variables defined in the collection by default.
However, if changing these variables is needed (e.g., to specify a different API), it is possible to set them from the command line using the `--env-var` parameter.

Below is an example where a request is ran but targeted to a specific API, with verbose output:

```console
newman run <path-to-postman-collection> --folder 'Get Collection with NFTs, paginated - first 200' --env-var scheme=http --env-var graph_hostname=localhost --env-var port=8080 --verbose
```

Instead of sending the request to the defined address of `https://dev-analytics-graph.nft.com:443/`, it will be sent to `http://localhost:8080/`.
