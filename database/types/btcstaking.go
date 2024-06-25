package types

type BtcStakingMintRow struct {
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

func NewBtcStakingMintRow(
	txHash string,
	height int64,
	signer string,
	receiverName string,
	receiverBtcAddress string,
	receiverEthAddress string,
	amount float64,
	btcStakingTxHash string,
) BtcStakingMintRow {
	return BtcStakingMintRow{
		TxHash:             txHash,
		Height:             height,
		Signer:             signer,
		ReceiverName:       receiverName,
		ReceiverBtcAddress: receiverBtcAddress,
		ReceiverEthAddress: receiverEthAddress,
		Amount:             amount,
		BtcStakingTxHash:   btcStakingTxHash,
	}
}

type BtcStakingBurnRow struct {
	ID               int64   `db:"id"`
	TxHash           string  `db:"tx_hash"`
	Height           int64   `db:"height"`
	Signer           string  `db:"signer"`
	TargetBtcAddress string  `db:"target_btc_address"`
	Amount           float64 `db:"amount"`
}

func NewBtcStakingBurnRow(
	txHash string,
	height int64,
	signer string,
	targetBtcAddress string,
	amount float64,
) BtcStakingBurnRow {
	return BtcStakingBurnRow{
		TxHash:           txHash,
		Height:           height,
		Signer:           signer,
		TargetBtcAddress: targetBtcAddress,
		Amount:           amount,
	}
}
