package customized

import (
	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/spf13/cobra"
)

func NewUtilsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "utils",
		Short: "Utilities that doesn't affect database",
	}

	cmd.AddCommand(
		newMissingBlocksCmd(parseConfig),
	)

	return cmd
}
