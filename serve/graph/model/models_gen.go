// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// Filters for querying Blocks within specified criteria related to their attributes.
type BlockFilter struct {
	// Minimum block height from which to start fetching Blocks, inclusive. If unspecified, there is no lower bound.
	FromHeight *int `json:"from_height,omitempty"`
	// Maximum block height up to which Blocks should be fetched, exclusive. If unspecified, there is no upper bound.
	ToHeight *int `json:"to_height,omitempty"`
	// Minimum timestamp from which to start fetching Blocks, inclusive. Blocks created at or after this time will be included.
	FromTime *time.Time `json:"from_time,omitempty"`
	// Maximum timestamp up to which to fetch Blocks, exclusive. Only Blocks created before this time are included.
	ToTime *time.Time `json:"to_time,omitempty"`
}

// Root Query type to fetch data about Blocks and Transactions based on filters or retrieve the latest block height.
type Query struct {
}

// Subscriptions provide a way for clients to receive real-time updates about Transactions and Blocks based on specified filter criteria.
// Subscribers will only receive updates for events occurring after the subscription is established.
type Subscription struct {
}

// Filters for querying Transactions within specified criteria related to their execution and placement within Blocks.
type TransactionFilter struct {
	// Minimum block height from which to start fetching Transactions, inclusive. Aids in scoping the search to recent Transactions.
	FromBlockHeight *int `json:"from_block_height,omitempty"`
	// Maximum block height up to which Transactions should be fetched, exclusive. Helps in limiting the search to older Transactions.
	ToBlockHeight *int `json:"to_block_height,omitempty"`
	// Minimum Transaction index from which to start fetching, inclusive. Facilitates ordering in Transaction queries.
	FromIndex *int `json:"from_index,omitempty"`
	// Maximum Transaction index up to which to fetch, exclusive. Ensures a limit on the ordering range for Transaction queries.
	ToIndex *int `json:"to_index,omitempty"`
	// Minimum `gas_wanted` value to filter Transactions by, inclusive. Filters Transactions based on the minimum computational effort declared.
	FromGasWanted *int `json:"from_gas_wanted,omitempty"`
	// Maximum `gas_wanted` value for filtering Transactions, exclusive. Limits Transactions based on the declared computational effort.
	ToGasWanted *int `json:"to_gas_wanted,omitempty"`
	// Minimum `gas_used` value to filter Transactions by, inclusive. Selects Transactions based on the minimum computational effort actually used.
	FromGasUsed *int `json:"from_gas_used,omitempty"`
	// Maximum `gas_used` value for filtering Transactions, exclusive. Refines selection based on the computational effort actually consumed.
	ToGasUsed *int `json:"to_gas_used,omitempty"`
}