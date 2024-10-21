package main

import (
	"context"
	"fmt"
	brokerPkg "github.com/kasasunil344/machineCoding/pubSub/broker"
	"github.com/kasasunil344/machineCoding/pubSub/cerror"
	"github.com/kasasunil344/machineCoding/pubSub/constants"
	"sync"
)

type PubSub struct {
	broker brokerPkg.Broker
}

func (ps *PubSub) Send(topic string, payload interface{}) cerror.CError {

	return ps.broker.AcknowledgeSubscribers(topic, payload)
}

func (ps *PubSub) Consume(topic string, name string) (interface{}, cerror.CError) {

	return ps.broker.GetPayload(context.Background(), name, topic)
}

func main() {
	// Initialize broker
	broker := brokerPkg.NewBroker()

	pubSub := PubSub{
		broker: broker,
	}

	// Register topic
	pubSub.broker.RegisterTopic(constants.Topic1)
	pubSub.broker.RegisterTopic(constants.Topic2)

	// Register subscriber
	pubSub.broker.Subscribe(constants.Topic1, constants.Subscriber1)
	pubSub.broker.Subscribe(constants.Topic1, constants.Subscriber2)

	pubSub.broker.Subscribe(constants.Topic2, constants.Subscriber1)
	pubSub.broker.Subscribe(constants.Topic2, constants.Subscriber2)

	// Publish and consume messages
	data := make(map[string]interface{})
	data["key1"] = "value1"
	data["key2"] = "value2"

	wg1 := sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		wg1.Add(1)
		go func(j int) {
			defer wg1.Done()
			err := pubSub.Send(fmt.Sprintf("topic%v", j), data)
			if err != nil {
				fmt.Println("Error publishing Data ", err)
			} else {
				fmt.Println("Published data Successfully!!")
			}
		}(i)
	}
	wg1.Wait()

	var wg2 sync.WaitGroup
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			wg2.Add(1)
			go func(x, y int) {
				defer wg2.Done()
				conData, err := pubSub.Consume(fmt.Sprintf("topic%v", x), fmt.Sprintf("subscriber%v", y))
				if err != nil {
					fmt.Println("Error Consuming Data ", err)
				} else {
					fmt.Println("Consumed data ", conData)
				}
			}(i, j)
		}
	}
	wg2.Wait()

}
