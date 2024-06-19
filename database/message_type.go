package database

import (
	"github.com/forbole/callisto/v4/types"
)

// SaveMessageType stores the given message type inside the database
func (db *Db) SaveMessageType(msg *types.MessageType) error {
	stmt := `
INSERT INTO message_type(type, module, label) 
VALUES ($1, $2, $3) 
ON CONFLICT (type) DO NOTHING`

	_, err := db.SQL.Exec(stmt, msg.Type, msg.Module, msg.Label)
	return err
}
