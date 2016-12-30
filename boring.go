package main

import (
	"fmt"
	"time"
)

func boring(c chan string) {
	for i := 0; ; i++ {
		s := fmt.Sprintf("boring: %v", i)
		c <- s
	}
}

func main() {
	c := make(chan string)
	go boring(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("[MAIN] %v\n", <-c)
		time.Sleep(time.Second)
	}
	fmt.Println("Um, like, this is _so_ gd boring")
}
