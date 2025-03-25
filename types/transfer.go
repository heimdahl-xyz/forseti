package types

import (
	"math/big"
)

type FTMessage struct {
	Timestamp    int64    `json:"timestamp" db:"timestamp"`
	FromAddress  string   `json:"from_address" db:"from_address"`
	FromOwner    string   `json:"from_owner,omitempty" db:"from_owner"`
	ToAddress    string   `json:"to_address" db:"to_address"`
	ToOwner      string   `json:"to_owner,omitempty" db:"to_owner"`
	Amount       *big.Int `json:"amount" db:"amount"`
	TokenAddress string   `json:"token_address" db:"token_address"`
	Symbol       string   `json:"symbol" db:"symbol"`
	Chain        string   `json:"chain" db:"chain"`
	Network      string   `json:"network" db:"network"`
	TxHash       string   `json:"tx_hash" db:"tx_hash"`
	Decimals     uint8    `json:"decimals" db:"decimals"`
	Position     uint64   `json:"position" db:"position"`
}
