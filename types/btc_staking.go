package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type BtcStakingMint struct {
	TxHash             string
	Height             int64
	Signer             string
	ReceiverName       string
	ReceiverBtcAddress string
	ReceiverEthAddress string
	Amount             sdk.Coin
	BtcStakingTxHash   string
}

func NewBtcStakingMint(
	txHash string,
	height int64,
	signer string,
	receiverName string,
	receiverBtcAddress string,
	receiverEthAddress string,
	amount sdk.Coin,
	btcStakingTxHash string,
) BtcStakingMint {
	return BtcStakingMint{
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

type BtcStakingBurn struct {
	TxHash           string
	Height           int64
	Signer           string
	TargetBtcAddress string
	Amount           sdk.Coin
}
