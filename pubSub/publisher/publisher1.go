package publisher

import (
	"github.com/kasasunil344/machineCoding/pubSub/broker"
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
)

type Publisher1 struct {
	broker broker.Broker
}

func New(bro broker.Broker) *Publisher1 {
	return &Publisher1{
		broker: bro,
	}
}

func (t1 *Publisher1) Send(topic string, payload interface{}) cerror.CError {
	// For now, I am sending payload in param,
	// ideally it GetPublishPayload() should be implemented by publisher.
	return t1.broker.AcknowledgeSubscribers(topic, payload)
}

func (t1 *Publisher1) GetPublishPayload() interface{} {
	data := make(map[string]interface{})
	data["key1"] = "value1"
	data["key2"] = "value2"

	return data
}
