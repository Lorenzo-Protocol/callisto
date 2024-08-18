package btcstaking_test

import (
	"encoding/json"
	"testing"

	btcstakingtypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestStakingRecordFromEvents(t *testing.T) {
	stakingRecord := btcstakingtypes.BTCStakingRecord{
		TxHash:          []byte("A1SlLHnKjNhVi8CiI4Qqqr2m2YCHrCL8wV9FwkBdggQ="),
		Amount:          100000000,
		MintToAddr:      []byte("cigaYCU50uMaESne59JsSWDI94c="),
		BtcReceiverName: "lorenzo",
		BtcReceiverAddr: "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq",
	}

	manager := sdk.NewEventManager()
	err := manager.EmitTypedEvent(btcstakingtypes.NewEventBTCStakingCreated(&stakingRecord))
	if err != nil {
		t.Fatal(err)
	}

	msg, _ := sdk.ParseTypedEvent(manager.Events().ToABCIEvents()[0])
	json.Unmarshal([]byte(msg.String()), &stakingRecord)

	t.Logf("Staking record: %v", stakingRecord)
}
