package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() { // receives channel c from func `boring`
	quit := make(chan string)
	c := boring("A", quit)

	num := rand.Intn(25)
	fmt.Printf("num times: %v\n", num)
	for i := num; i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- ""
	fmt.Printf("A quit and says, %q\n", <-quit)

}

func boring(name string, quit chan string) <-chan string { // returns receive-only channel
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", name, i):
				time.Sleep(time.Duration(rand.Intn(1.25e3)) * time.Millisecond)
			case <-quit:
				cleanup()
				quit <- "Bye bye!"
				return
			}
		}
	}()

	return c
}

func cleanup() {
	fmt.Println("cleaning up!")
	time.Sleep(1 * time.Second)
}
