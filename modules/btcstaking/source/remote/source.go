package remote

import (
	"github.com/forbole/juno/v5/node/remote"

	btcstakingtypes "github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	btcstakingsource "github.com/forbole/callisto/v4/modules/btcstaking/source"
)

var (
	_ btcstakingsource.Source = &Source{}
)

// Source implements btcstakingsouce.Source using a remote node
type Source struct {
	*remote.Source
	querier btcstakingtypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, querier btcstakingtypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// StakingRecord implements btcstakingsource.Source
func (s Source) StakingRecord(height int64, hash []byte) (*btcstakingtypes.BTCStakingRecord, error) {
	res, err := s.querier.StakingRecord(remote.GetHeightRequestContext(s.Ctx, height),
		&btcstakingtypes.QueryStakingRecordRequest{TxHash: hash})
	if err != nil {
		return nil, err
	}

	return res.Record, nil
}
