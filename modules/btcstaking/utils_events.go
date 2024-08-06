package btcstaking

import (
	"fmt"

	btcstakingtypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	eventsutil "github.com/forbole/callisto/v4/utils/events"
)

const (
	EventTypeBtcStakingCreated = "lorenzo.btcstaking.v1.EventBTCStakingCreated"
	EventAttributeRecord       = "record"
)

// StakingRecordFromEvents returns the staking record from the given events
func (m *Module) StakingRecordFromEvents(events sdk.StringEvents) (*btcstakingtypes.BTCStakingRecord, error) {
	record, ok := eventsutil.FindEventByType(events, EventTypeBtcStakingCreated)
	if ok {
		return m.ParseStakingRecord(record)
	}

	return nil, fmt.Errorf("no staking record found")
}

// ParseStakingRecord returns the staking record from the given string
func (m *Module) ParseStakingRecord(record sdk.StringEvent) (*btcstakingtypes.BTCStakingRecord, error) {
	attributes := make([]abci.EventAttribute, 0)
	for _, attribute := range record.Attributes {
		attributes = append(attributes, attribute.ToKVPair())
	}

	event := abci.Event{
		Type:       record.Type,
		Attributes: attributes,
	}

	msg, err := sdk.ParseTypedEvent(event)
	if err != nil {
		return nil, fmt.Errorf("failed to parse staking record: %s", err)
	}

	eventMsg, ok := msg.(*btcstakingtypes.EventBTCStakingCreated)
	if !ok {
		return nil, fmt.Errorf("invalid event type: %T", msg)
	}

	return eventMsg.Record, nil
}
