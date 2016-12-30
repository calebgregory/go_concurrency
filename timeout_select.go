package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() { // receives channel c from func `boring`
	c := boring("A")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You are toooooo slow")
			return
		}

	}
}

func boring(name string) <-chan string { // returns receive-only channel
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			m := fmt.Sprintf("%s: %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1.25e3)) * time.Millisecond)
			c <- m
		}
	}()

	return c
}
