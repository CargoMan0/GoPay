package kafka

import (
	"fmt"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/core"
)

func fromDomainToKafka(eventName core.EventName) string {
	switch eventName {
	case core.AccountBalanceCreated:
		return TopicAccountBalanceCreated
	case core.AccountBalanceCreationFailed:
		return TopicAccountBalanceCreationFailed
	default:
		panic(fmt.Sprintf("unknown event from core: (%q)", eventName))
	}
}
