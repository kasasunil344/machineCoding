package main

import (
	"fmt"
	"sync"
)

type Agent struct {
	mu     sync.Mutex
	subs   map[string][]chan string
	closed bool
}

func NewAgent() *Agent {
	return &Agent{
		subs: make(map[string][]chan string),
	}
}

func (ag *Agent) Publish(wg1 *sync.WaitGroup, topic string, msg string) {
	if _, ok := ag.subs[topic]; !ok {
		fmt.Println("Invalid topic!!!")
		return
	}

	ag.mu.Lock()
	defer func() {
		ag.mu.Unlock()
		wg1.Done()
	}()

	if ag.closed {
		fmt.Println("Publisher -> Agent is closed!!!!")
		return
	}

	for _, ch := range ag.subs[topic] {
		ch <- msg
	}
}

func (ag *Agent) Subscribe(topic string) chan string {
	ag.mu.Lock()
	defer ag.mu.Unlock()

	if ag.closed {
		fmt.Println("Subscriber -> Agent is closed!!!!")
		return nil
	}

	ch := make(chan string)
	ag.subs[topic] = append(ag.subs[topic], ch)
	return ch
}

func (ag *Agent) close() {
	ag.mu.Lock()
	defer ag.mu.Unlock()

	if ag.closed {
		fmt.Println("Agent is closed!!!!")
		return
	}

	ag.closed = true
	for _, ch := range ag.subs {
		for _, sub := range ch {
			close(sub)
		}
	}
}

func main() {
	agent := NewAgent()

	ch := agent.Subscribe("topic1")

	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	go agent.Publish(&wg1, "topic1", "second pub/sub implementation")

	fmt.Println(<-ch)

	wg1.Wait()

	agent.close()

}
