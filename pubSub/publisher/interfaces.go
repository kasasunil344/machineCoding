package publisher

import (
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
)

type Publisher interface {
	Send(topic string, payload interface{}) cerror.CError
}
