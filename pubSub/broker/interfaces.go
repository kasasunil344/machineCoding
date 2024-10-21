package broker

import (
	"context"
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
)

/*
--------NOTE:----------
If we are using DB for storing topics, subscriber data and payloads we should use individual interfaces,
but to simulate it without using DB we need to keep all interfaces into one because we will get new broker
initialized for every instance.


type TopicRegistration interface {
	RegisterTopic(topic string) cerror.CError
}

type Subscribe interface {
	Subscribe(topic string, subscriberNM string) cerror.CError
}

type PublisherInterfaces interface {
	AcknowledgeSubscribers(topic string, payload interface{}) cerror.CError
}

type ConsumerInterfaces interface {
	GetPayload(ctx context.Context,subscriberName string, topic string) (interface{}, cerror.CError)
}

*/

type Broker interface {
	RegisterTopic(topic string) cerror.CError
	Subscribe(topic string, subscriberNM string) cerror.CError
	AcknowledgeSubscribers(topic string, payload interface{}) cerror.CError
	GetPayload(ctx context.Context, subscriberName string, topic string) (interface{}, cerror.CError)
}
