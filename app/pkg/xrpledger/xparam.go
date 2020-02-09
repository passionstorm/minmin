package main

type Xparam struct {
	Method string `json:"method"`
	Params []interface{} `json:"params"`
}

/**
Use the wallet_propose method to generate a key pair and XRP Ledger address.
This command only generates key and address values,
and does not affect the XRP Ledger itself in any way.
To become a funded address stored in the
*/
type WalletProposeParam struct {
	Passphrase string `json:"passphrase,omitempty"`
	KeyType	string `json:"key_type,omitempty"`
	Seed string `json:"seed,omitempty"`
	SeedHex string `json:"seed_hex,omitempty"`
}