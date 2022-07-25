# Graph Pagination Verify

## Description

This tool can be used to verify that the Graph API pagination is working correctly.
First, the tool connects to the database and retrieves the list of NFTs in the specified collection.
The same collection is then retrieved from the Graph API, requesting NFTs in batches of the specified size.
A check is made to ensure that NFTs are returned in the correct (sorted) order, as well as that the cursors have the correct value.
Finally, when all pages are processed, it is checked whether the Graph API indicated that there are no more pages to be processed.

If needed, it's possible to write all Graph API responses to a file so they can be closely examined.

## Usage

```
  -c, --collection-id string   collection ID (default "612ecc22-36ef-4ef7-bb0b-5b864b85d089")
  -d, --database string        database address
      --dump-file string       dump file to write response data to
  -g, --graph-api string       Graph API endpoint (default "http://127.0.0.1:8080/graphql")
  -l, --log-level string       log level (default "info")
  -p, --page-size uint         number of items to request in a single page (default 250)
```
