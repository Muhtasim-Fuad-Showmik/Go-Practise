package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	dones := make([]chan bool, 4)

	dones[0] = make(chan bool)
	dones[1] = make(chan bool)
	dones[2] = make(chan bool)
	dones[3] = make(chan bool)

	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])

	// You can print the channel output if you want
	// fmt.Println(<-done)

	for _, done := range dones {
		<-done
	}
}
