package broker

import (
	"context"
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
	"sync"
)

type MessageQ struct {
	topicSubsData map[string][]subscriber
	mutex         sync.Mutex
}

type subscriber struct {
	name string
	data chan interface{}
}

func NewBroker() Broker {
	return &MessageQ{
		topicSubsData: make(map[string][]subscriber),
	}
}

func (m *MessageQ) RegisterTopic(topic string) cerror.CError {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.topicSubsData[topic] = []subscriber{}
	return nil
}

func (m *MessageQ) Subscribe(topic string, subscriberNM string) cerror.CError {
	if _, ok := m.topicSubsData[topic]; ok {
		m.topicSubsData[topic] = append(m.topicSubsData[topic], subscriber{
			name: subscriberNM,
			data: make(chan interface{}),
		})
	} else {
		return cerror.New("ERROR_INVALID_TOPIC_NAME")
	}
	return nil
}

func (m *MessageQ) AcknowledgeSubscribers(topic string, payload interface{}) cerror.CError {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	subs, err := m.getSubscribers(topic)
	if err != nil {
		return err
	}

	for _, val := range subs {
		go func(v subscriber) {
			v.data <- payload
		}(val)
	}

	return nil
}

func (m *MessageQ) getSubscribers(topic string) ([]subscriber, cerror.CError) {
	if _, ok := m.topicSubsData[topic]; !ok {
		return []subscriber{}, cerror.New("ERROR_INVALID_TOPIC_NAME")
	}

	return m.topicSubsData[topic], nil
}

func (m *MessageQ) GetPayload(ctx context.Context, subscriberName string, topic string) (interface{}, cerror.CError) {
	// Find subscriber by combination topic and subscriberName
	sub, err := m.findSubscriberByTopicAndName(topic, subscriberName)
	if err != nil {
		return nil, err
	}

	return <-sub.data, nil
}

func (m *MessageQ) findSubscriberByTopicAndName(topic string, name string) (*subscriber, cerror.CError) {
	subscribers, err := m.getSubscribers(topic)
	if err != nil {
		return nil, err
	}

	for _, sub := range subscribers {
		if sub.name == name {
			return &sub, nil
		}
	}

	return nil, cerror.New("ERROR_SUBSCRIBER_NOT_FOUND")
}
