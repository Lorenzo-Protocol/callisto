package btcstaking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	btcstakingsource "github.com/forbole/callisto/v4/modules/btcstaking/source"
	"github.com/forbole/juno/v5/modules"

	"github.com/forbole/callisto/v4/database"
)

var (
	_ modules.Module = &Module{}
)

// Module represent x/gov module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source btcstakingsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source btcstakingsource.Source,
	cdc codec.Codec,
	db *database.Db,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements
func (m *Module) Name() string { return "btcstaking" }
