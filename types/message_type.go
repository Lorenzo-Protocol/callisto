package types

// MessageType represents the data of a single message type
type MessageType struct {
	Type   string
	Module string
	Label  string
}

// NewMessageType allows to build a new MessageType instance
func NewMessageType(msgType string, module string, label string) *MessageType {
	return &MessageType{
		Type:   msgType,
		Module: module,
		Label:  label,
	}
}
