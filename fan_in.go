package main

import (
	"fmt"
	"time"
)

func main() { // receives channel c from func `boring`
	c := fanIn(boring("A"), boring("C"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		time.Sleep(time.Second / 2)
	}
	fmt.Println("omg tihs is s0e broing gime outahere")
}

func boring(name string) <-chan string { // returns receive-only channel
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			m := fmt.Sprintf("%v: %v", name, i)
			c <- m
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
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
