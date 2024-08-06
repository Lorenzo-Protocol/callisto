package btcstaking

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	btcstakingtypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *btcstakingtypes.MsgCreateBTCStaking:
		return m.handleCreateBTCStaking(tx, cosmosMsg.Signer, tx.Logs[index].Events)
	case *btcstakingtypes.MsgBurnRequest:
		return m.handleBurnRequest(tx, cosmosMsg)
	}
	return nil
}

func (m *Module) handleCreateBTCStaking(tx *juno.Tx, signer string, events sdk.StringEvents) error {
	record, err := m.StakingRecordFromEvents(events)
	if err != nil {
		return fmt.Errorf("error while getting btc staking record: %s", err)
	}

	err = m.db.SaveBtcStakingMint(tx.Height, tx.TxHash, signer, record)
	if err != nil {
		return fmt.Errorf("error while saving btc staking mint: %s", err)
	}

	return nil
}

func (m *Module) handleBurnRequest(tx *juno.Tx, msg *btcstakingtypes.MsgBurnRequest) error {
	err := m.db.SaveBtcStakingBurn(
		tx.Height,
		tx.TxHash,
		msg.Signer,
		msg.BtcTargetAddress,
		msg.Amount.Int64(),
	)
	if err != nil {
		return fmt.Errorf("error while saving btc staking burn: %s", err)
	}

	return nil
}
