package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	// 0. A function in a goroutine runs "in the background", meaning that
	//    the runtime does not wait for `boring(c)` to finish executing to resume
	//    executing the next line of code (*)-,
	go boring(c)             //               |
	for i := 0; i < 5; i++ { // <-------------'
		// 1. When the main function executes `<-c`, it _waits_ for a value to be sent.
		fmt.Printf("[MAIN] %v\n", <-c) // [RECEIVER]
		time.Sleep(time.Second)
	}
	fmt.Println("Um, like, this is _so_ gd boring")
}

func boring(c chan string) {
	for i := 0; ; i++ {
		s := fmt.Sprintf("boring: %v", i)
		// 2. Similarly, when a function in a goroutine executes `c <- value`, it
		//    waits for the receiver to be ready.
		c <- s // [SENDER]
	}
}

// A sender and receiver must both be ready to perform their part in the
// communication. Otherwise, the runtime waits until they are.

// Thus channels both communicate and synchronize.
