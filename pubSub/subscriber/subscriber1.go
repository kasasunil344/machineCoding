package subscriber

import (
	"context"
	"github.com/kasasunil344/machineCoding/pubSub/broker"
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
)

type Subscriber1 struct {
	broker broker.Broker
}

func (s1 *Subscriber1) Consume(topic string, name string) (interface{}, cerror.CError) {
	return s1.broker.GetPayload(context.Background(), name, topic)
}
