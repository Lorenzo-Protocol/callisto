package customized

import (
	"fmt"

	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/forbole/juno/v5/types/config"
	"github.com/spf13/cobra"
)

func newMissingBlocksCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "missing-blocks",
		Short: "Get all the missing heights in the database by end of the latest height",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			dbLastHeight, err := parseCtx.Database.GetLastBlockHeight()
			if err != nil {
				return fmt.Errorf("error while getting DB last block height: %s", err)
			}

			missingBlocks := parseCtx.Database.GetMissingHeights(1, dbLastHeight)
			fmt.Println("Missing blocks: ", len(missingBlocks))
			fmt.Println("Missing blocks: ", missingBlocks)
			return nil
		},
	}

	return cmd
}
