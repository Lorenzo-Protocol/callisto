package database

import (
	"encoding/hex"
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func (db *Db) SaveBtcStakingMint(height int64, txHash string, signer string, record *types.BTCStakingRecord) error {
	stmt := `
INSERT INTO btc_staking_mint(tx_hash, height, signer, receiver_name, receiver_btc_address, receiver_eth_address, amount, btc_staking_tx_hash)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

	_, err := db.SQL.Exec(stmt,
		txHash,
		height,
		signer,
		record.AgentName,
		record.AgentBtcAddr,
		"0x"+hex.EncodeToString(record.ReceiverAddr),
		record.Amount,
		(chainhash.Hash)(record.TxHash).String(),
	)
	if err != nil {
		return fmt.Errorf("error while saving btc staking mint: %s", err)
	}

	return nil
}

func (db *Db) SaveBtcStakingBurn(height int64, txHash, signer, targetAddress string, amount int64) error {
	stmt := `
INSERT INTO btc_staking_burn(tx_hash, height, signer, target_btc_address, amount)
VALUES ($1, $2, $3, $4, $5)
`

	_, err := db.SQL.Exec(stmt,
		txHash,
		height,
		signer,
		targetAddress,
		amount,
	)
	if err != nil {
		return fmt.Errorf("error while saving btc staking burn: %s", err)
	}

	return nil
}
