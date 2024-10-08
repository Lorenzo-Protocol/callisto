package customized

import (
	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/spf13/cobra"
)

func NewUtilsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "utils",
		Short: "Customized utilities",
	}

	cmd.AddCommand(
		newMissingBlocksCmd(parseConfig),
		newMissingTransactionsCmd(parseConfig),
	)

	return cmd
}
