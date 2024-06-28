package customized

import (
	"sort"
	"strconv"
	"strings"

	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/forbole/juno/v5/parser"
	"github.com/forbole/juno/v5/types/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newTransactionsCmd returns a Cobra command that allows to fix missing or incomplete transactions in database

func newMissingTransactionsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handle-missing-txs",
		Short: "Parse missing or incomplete transactions",
		Args:  cobra.ExactArgs(1),
		Long: `
Maintainer need to query heights missing transactions in database and parse them as a string. 

Example:
	SELECT STRING_AGG(CAST(b.height AS TEXT), ', ')
	FROM block b
	JOIN (
		SELECT height, COUNT(*) as tx_count
		FROM transaction
		GROUP BY height
	) t ON b.height = t.height
	WHERE b.num_txs > 0 AND b.num_txs != t.tx_count;

callisto parse utils handle-missing-txs "1, 2, 3, 4, 5"
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			workerCtx := parser.NewContext(parseCtx.EncodingConfig, parseCtx.Node, parseCtx.Database, parseCtx.Logger, parseCtx.Modules)
			worker := parser.NewWorker(workerCtx, nil, 0)

			// parse height strings to heights array
			strs := strings.Split(strings.TrimSpace(args[0]), ",")
			heights := make([]int64, len(strs))
			for i, s := range strs {
				height, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					return err
				}
				heights[i] = int64(height)
			}

			sort.Slice(heights, func(i, j int) bool { return heights[i] < heights[j] })
			log.Info().Msgf("Processing heights: %v", heights)

			unhandles := make([]int64, 0)
			for _, h := range heights {
				log.Info().Int64("height", h).Msg("processing transactions...")
				err = worker.ProcessTransactions(h)
				if err != nil {
					log.Error().Msgf("error while re-fetching transactions of height %d: %s", h, err)
					unhandles = append(unhandles, h)
				}
			}
			log.Info().Msgf("Unprocessed heights: %v", unhandles)

			return nil
		},
	}
	return cmd
}
