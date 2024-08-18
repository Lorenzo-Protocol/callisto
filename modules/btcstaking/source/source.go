package source

import btcstakingtypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"

type Source interface {
	StakingRecord(int64, []byte) (*btcstakingtypes.BTCStakingRecord, error)
}
