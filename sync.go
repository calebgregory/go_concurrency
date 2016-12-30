package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() { // receives channel c from func `boring`
	c := fanIn(boring("A"), boring("C"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("omg tihs is s0e broing gime outahere")
}

func boring(name string) <-chan Message { // returns receive-only channel
	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			m := Message{fmt.Sprintf("%s: %d", name, i), waitForIt}
			c <- m
			time.Sleep(time.Duration(200) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
