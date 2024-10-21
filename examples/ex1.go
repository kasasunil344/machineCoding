package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var Ch = make(chan string)

func main() {

	for i := range 5 {
		wg.Add(1)
		go publishMsg(fmt.Sprintf("count %v", i))
		wg.Add(1)
		go receiveMsg()
	}

	// Wait for the goroutines to finish
	wg.Wait()
}

func publishMsg(msg string) {
	defer wg.Done() // Signal that this goroutine is done
	Ch <- msg
}

func receiveMsg() {
	defer wg.Done() // Signal that this goroutine is done
	msg := <-Ch
	fmt.Println("Message received : ", msg)
}
