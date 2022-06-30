// Package query describes the GraphQL queries requesting Collection, Marketplace or NFT
// data from the GraphQL server, in accordance with the GraphQL schema defined in
// graph/schema/schema.graphql.
//
// Queries are parsed according to the provided configuration (path to each of the wanted fields)
// and allow easy inspection to determine which fields were requested. Not all possible fields
// for a type are modeled, but only those that carry expensive overhead in retrieving or calculating,
// such as trait or rarity information, or any aggregated fields retrieved from the Aggregation API.

package query
