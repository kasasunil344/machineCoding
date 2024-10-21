package subscriber

import (
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
)

type Subscriber interface {
	Consume(topic string, name string) (interface{}, cerror.CError)
}
