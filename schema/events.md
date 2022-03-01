# Events API

- Based on REST API with paging/limit & filter parameters
- Type is not a filter parameter, one endpoint per type
  - avoids having to deal with polymorphism of events
  - mint, burn, transfer, sale, ...
- events.nft.com/mints?limit=1000&before=12345&collection=...

If someone goes to an NFT and wants to see all events of that NFT, the frontend has to query multiple endpoints and merge the events list.
The upside is, however, that it will be easy for them to filter subsequently.