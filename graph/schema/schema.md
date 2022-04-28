# Schema Types

<details>

  <summary><strong>Table of Contents</strong></summary>

  * [Query](#query)
  * [Objects](#objects)
    * [Chain](#chain)
    * [Collection](#collection)
    * [Marketplace](#marketplace)
    * [NFT](#nft)
	* [Trait](#trait)
  * [Inputs](#inputs)
    * [CollectionOrder](#collectionorder)
    * [NFTOrder](#nftorder)
  * [Enums](#enums)
    * [CollectionOrderField](#collectionorderfield)
    * [NFTOrderField](#nftorderfield)
    * [OrderDirection](#orderdirection)
  * [Scalars](#scalars)
    * [Address](#address)
    * [Boolean](#boolean)
    * [DateTime](#datetime)
    * [Float](#float)
    * [ID](#id)
    * [String](#string)

</details>

## Query

The query root of NFT.com GraphQL interface.

<table>
	<thead>
		<tr>
			<th align="left">Field</th>
			<th align="right">Argument</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>chain</strong></td>
			<td valign="top"><a href="#chain">Chain</a></td>
			<td>Get a single chain.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">id</td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>ID of the Chain.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>chains</strong></td>
			<td valign="top">[<a href="#chain">Chain</a>!]</td>
			<td>List chains.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>nft</strong></td>
			<td valign="top"><a href="#nft">NFT</a></td>
			<td>Get a single NFT.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">id</td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>ID of the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>nftByTokenID</strong></td>
			<td valign="top"><a href="#nft">NFT</a></td>
			<td>Get a single NFT by its token ID.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">chainID</td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>Chain ID.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">contract</td>
			<td valign="top"><a href="#address">Address</a>!</td>
			<td>ID of the smart contract.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">tokenID</td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Token ID of the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>nfts</strong></td>
			<td valign="top">[<a href="#nft">NFT</a>!]</td>
			<td>Lookup NFTs based on specified criteria.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">owner</td>
			<td valign="top"><a href="#address">Address</a></td>
			<td>Owner of the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">collection</td>
			<td valign="top"><a href="#id">ID</a></td>
			<td>ID of the collection the NFT is part of.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">rarityMax</td>
			<td valign="top"><a href="#float">Float</a></td>
			<td>Maximum rarity value.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">orderBy</td>
			<td valign="top"><a href="#nftorder">NFTOrder</a></td>
			<td>Ordering options for the returned NFTs.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>collection</strong></td>
			<td valign="top"><a href="#collection">Collection</a></td>
			<td>Get a single collection.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">id</td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>ID of the collection.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>collectionByAddress</strong></td>
			<td valign="top"><a href="#collection">Collection</a></td>
			<td>Get a single collection by address.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">chainID</td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>Chain ID.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">contract</td>
			<td valign="top"><a href="#address">Address</a>!</td>
			<td>Address of the smart contract.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>collections</strong></td>
			<td valign="top">[<a href="#collection">Collection</a>!]</td>
			<td>Lookup collections based on specified criteria.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">chain</td>
			<td valign="top"><a href="#id">ID</a></td>
			<td>ID of the chain that the collection is on.</td>
		</tr>
		<tr>
			<td colspan="2" align="right" valign="top">orderBy</td>
			<td valign="top"><a href="#collectionorder">CollectionOrder</a></td>
			<td>Ordering options for the returned collections.</td>
		</tr>
	</tbody>
</table>

## Objects

### Chain

Chain represents the chain and its networks.

<table>
	<thead>
		<tr>
			<th align="left">Field</th>
			<th align="right">Argument</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>id</strong></td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>Chain ID.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>name</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Name of the chain, e.g. `Ethereum`.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>description</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Description of the chain.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>marketplaces</strong></td>
			<td valign="top">[<a href="#marketplace">Marketplace</a>!]</td>
			<td>Marketplaces on this chain.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>collections</strong></td>
			<td valign="top">[<a href="#collection">Collection</a>!]</td>
			<td>Collections found on this chain.</td>
		</tr>
	</tbody>
</table>

### Collection

Collection represents a group of NFTs that share the same smart contract.

<table>
	<thead>
		<tr>
			<th align="left">Field</th>
			<th align="right">Argument</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>id</strong></td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>Collection ID.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>name</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Collection name, e.g. `CryptoKitties`.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>description</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Description of the collection.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>address</strong></td>
			<td valign="top"><a href="#address">Address</a>!</td>
			<td>Address of the smart-contract.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>website</strong></td>
			<td valign="top"><a href="#string">String</a></td>
			<td>Collection website.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>image_url</strong></td>
			<td valign="top"><a href="#string">String</a></td>
			<td>URL of an image for the collection.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>chain</strong></td>
			<td valign="top"><a href="#chain">Chain</a>!</td>
			<td>Chain on which collection resides on.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>marketplaces</strong></td>
			<td valign="top">[<a href="#marketplace">Marketplace</a>!]</td>
			<td>Marketplaces this collection is on.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>nfts</strong></td>
			<td valign="top">[<a href="#nft">NFT</a>!]</td>
			<td>List of NFTs that are part of this collection.</td>
		</tr>
	</tbody>
</table>

### Marketplace

Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).

<table>
	<thead>
		<tr>
			<th align="left">Field</th>
			<th align="right">Argument</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>id</strong></td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>Marketplace ID.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>name</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Marketplace name, e.g. `Opensea`.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>description</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Description of the marketplace.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>website</strong></td>
			<td valign="top"><a href="#string">String</a></td>
			<td>Marketplace website.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>chains</strong></td>
			<td valign="top">[<a href="#chain">Chain</a>!]!</td>
			<td>Chains the marketplace operates on.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>collections</strong></td>
			<td valign="top">[<a href="#collection">Collection</a>!]</td>
			<td>Collections on this marketplace.</td>
		</tr>
	</tbody>
</table>

### NFT

NFT represents a single Non-Fungible Token.

<table>
	<thead>
		<tr>
			<th align="left">Field</th>
			<th align="right">Argument</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>id</strong></td>
			<td valign="top"><a href="#id">ID</a>!</td>
			<td>NFT ID.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>tokenID</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Token ID, as found on the blockchain.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>name</strong></td>
			<td valign="top"><a href="#string">String</a></td>
			<td>Name of the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>image_url</strong></td>
			<td valign="top"><a href="#string">String</a></td>
			<td>URL of an image for the NFT.</td>
      <tr>
        <td colspan="2" valign="top"><strong>uri</strong></td>
        <td valign="top"><a href="#string">String</a>!</td>
        <td>URI directing to e.g. a JSON file with token metadata.</td>
      </tr>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>description</strong></td>
			<td valign="top"><a href="#string">String</a></td>
			<td>Description of the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>owner</strong></td>
			<td valign="top"><a href="#address">Address</a>!</td>
			<td>Address of the account that owns the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>rarity</strong></td>
			<td valign="top"><a href="#float">Float</a>!</td>
			<td>Rarity score for the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>traits</strong></td>
			<td valign="top">[<a href="#trait">Trait</a>!]</td>
			<td>Traits contains a list of attributes of the NFT.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>collection</strong></td>
			<td valign="top"><a href="#collection">Collection</a>!</td>
			<td>Collection this NFT is part of.</td>
		</tr>
	</tbody>
</table>

### Trait

Trait represents a single NFT trait.

<table>
	<thead>
		<tr>
			<th align="left">Field</th>
			<th align="right">Argument</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>type</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Trait type.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>value</strong></td>
			<td valign="top"><a href="#string">String</a>!</td>
			<td>Trait value.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>rarity</strong></td>
			<td valign="top"><a href="#float">Float</a>!</td>
			<td>Trait rarity represents the ratio of NFTs in a collection with this specific trait.</td>
		</tr>
	</tbody>
</table>

## Inputs

### CollectionOrder

Ordering options for collections.

<table>
	<thead>
		<tr>
			<th colspan="2" align="left">Field</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>field</strong></td>
			<td valign="top"><a href="#collectionorderfield">CollectionOrderField</a>!</td>
			<td>Field by which collections should be sorted.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>direction</strong></td>
			<td valign="top"><a href="#orderdirection">OrderDirection</a>!</td>
			<td>Direction in which collections should be sorted.</td>
		</tr>
	</tbody>
</table>

### NFTOrder

Ordering options for NFTs.

<table>
	<thead>
		<tr>
			<th colspan="2" align="left">Field</th>
			<th align="left">Type</th>
			<th align="left">Description</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td colspan="2" valign="top"><strong>field</strong></td>
			<td valign="top"><a href="#nftorderfield">NFTOrderField</a>!</td>
			<td>Field by which NFTs should be sorted by.</td>
		</tr>
		<tr>
			<td colspan="2" valign="top"><strong>direction</strong></td>
			<td valign="top"><a href="#orderdirection">OrderDirection</a>!</td>
			<td>Direction in which NFTs should be sorted.</td>
		</tr>
	</tbody>
</table>

## Enums

### CollectionOrderField

Properties by which collections can be ordered.

<table>
	<thead>
		<th align="left">Value</th>
		<th align="left">Description</th>
	</thead>
	<tbody>
		<tr>
			<td valign="top"><strong>CREATION_TIME</strong></td>
			<td>Order by creation time.</td>
		</tr>
		<tr>
			<td valign="top"><strong>MARKET_CAP</strong></td>
			<td>Order by market cap.</td>
		</tr>
		<tr>
			<td valign="top"><strong>TOTAL_VOLUME</strong></td>
			<td>Order by total volume.</td>
		</tr>
		<tr>
			<td valign="top"><strong>BIGGEST_GAINS</strong></td>
			<td>Order by biggest gains.</td>
		</tr>
		<tr>
			<td valign="top"><strong>BIGGEST_LOSSES</strong></td>
			<td>Order by biggest losses.</td>
		</tr>
		<tr>
			<td valign="top"><strong>DAILY_VOLUME</strong></td>
			<td>Order by daily volume.</td>
		</tr>
	</tbody>
</table>

### NFTOrderField

Properties by which NFTs could be ordered by.

<table>
	<thead>
		<th align="left">Value</th>
		<th align="left">Description</th>
	</thead>
	<tbody>
		<tr>
			<td valign="top"><strong>CREATION_TIME</strong></td>
			<td>Order by creation time.</td>
		</tr>
		<tr>
			<td valign="top"><strong>RARITY</strong></td>
			<td>Order by rarity.</td>
		</tr>
		<tr>
			<td valign="top"><strong>VALUE</strong></td>
			<td>Order by value.</td>
		</tr>
	</tbody>
</table>

### OrderDirection

Available options for the `orderBy` direction argument.

<table>
	<thead>
		<th align="left">Value</th>
		<th align="left">Description</th>
	</thead>
	<tbody>
		<tr>
			<td valign="top"><strong>ASC</strong></td>
			<td>Specifies an ascending order for a given `orderBy` argument.</td>
		</tr>
		<tr>
			<td valign="top"><strong>DESC</strong></td>
			<td>Specifies a decending order for a given `orderBy` argument.</td>
		</tr>
	</tbody>
</table>

## Scalars

### Address

A string representing an address (e.g. an Ethereum address).

### Boolean

The `Boolean` scalar type represents `true` or `false`.

### DateTime

An ISO-8601 encoded UTC date string, for example `2022-02-21T10:57:54Z`.

### Float

The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](https://en.wikipedia.org/wiki/IEEE_floating_point).

### ID

The `ID` scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as `"4"`) or integer (such as `4`) input value will be accepted as an ID.

### String

The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.

