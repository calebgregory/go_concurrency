package main

import (
	"fmt"
	"time"
)

func main() { // receives channel c from func `boring`
	c := boring("boring!")

	for i := 0; i < 5; i++ {
		fmt.Printf("u se: %q\n", <-c)
		time.Sleep(time.Second / 2)
	}
	fmt.Println("omg tihs is s0e broing gime outahere")
}

func boring(s string) <-chan string { // returns receive-only channel
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			m := fmt.Sprintf("%v %v", s, i)
			c <- m
		}
	}()

	return c
}
