package outbox

import "github.com/CargoMan0/GoPay/src/accountbalances/internal/core"

type Message struct {
	Name    core.EventName
	Status  MessageStatus
	Payload []byte
}

type MessageStatus uint8

const (
	MessageStatusPending MessageStatus = iota + 1
	MessageStatusSent
)
