# Postman collections

A single Postman collection file provides requests for each of the Analytic APIs: Graph, Events and Aggregation. These sample requests help showcase the different use case behind each API. While some input parameters are specified directly in the request, a handful of other variables are configurable on the Collection level. 

Note the env.json file is only used for the newman script below but it is not necessary to work with the above Collection. 

## Running tests

It is possible to run the collections and associated tests from the command line using a tool like [Newman](https://learning.postman.com/docs/running-collections/using-newman-cli/command-line-integration-with-newman/).

First Newman should be installed using the following command:

```console
$ npm install -g newman
```

Collections and tests checking the functionality of the `/transfers/` endpoint of the Events API can now be ran using:

```console
newman run events/transfers_collection.json -e env.json
```

Note that the targeted API should be running and reachable according to the provided environment.

### Environment

The `env.json` file describes the environment variables used by the Postman collections.
For example, the content of the `env.json` file might contain something like this:


```json
{
        "id": "3ba39c7b-49d2-4ad3-880b-c501ed7f043e",
        "values": [
                {
                        "key": "scheme",
                        "value": "http",
                        "type": "default",
                        "enabled": true
                },
                {
                        "key": "graph_hostname",
                        "value": "localhost",
                        "type": "default",
                        "enabled": true
                },
                                {
                        "key": "event_hostname",
                        "value": "localhost",
                        "type": "default",
                        "enabled": true
                },
                                {
                        "key": "aggregation_hostname",
                        "value": "localhost",
                        "type": "default",
                        "enabled": true
                },
                {
                        "key": "port",
                        "value": "8080",
                        "type": "default",
                        "enabled": true
                }
        ],
        "name": "Globals",
        "_postman_variable_scope": "globals",
        "_postman_exported_at": "2022-06-16T15:32:01.351Z",
        "_postman_exported_using": "Postman/9.21.2-220607-0647"
}
```

Requests in the Postman collections reference these variables in requests, by specifying the HTTP address of the API as e.g. `{{scheme}}://{{*_hostname}}:{{port}}/transfers/`.
Each of the variables - `scheme`, `*_hostname` and `port` are loaded from the provided environment file.
These variables should be set to the appropriate values for the API being tested.