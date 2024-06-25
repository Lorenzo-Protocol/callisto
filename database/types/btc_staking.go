package types

type BtcStakingMint struct {
	ID                 int64   `db:"id"`
	TxHash             string  `db:"tx_hash"`
	Height             int64   `db:"height"`
	Signer             string  `db:"signer"`
	ReceiverName       string  `db:"receiver_name"`
	ReceiverBtcAddress string  `db:"receiver_btc_address"`
	ReceiverEthAddress string  `db:"receiver_eth_address"`
	Amount             float64 `db:"amount"`
	BtcStakingTxHash   string  `db:"btc_staking_tx_hash"`
}

type BtcStakingBurn struct {
	ID               int64   `db:"id"`
	TxHash           string  `db:"tx_hash"`
	Height           int64   `db:"height"`
	Signer           string  `db:"signer"`
	TargetBtcAddress string  `db:"target_btc_address"`
	Amount           float64 `db:"amount"`
}
