package main

import "fmt"

func f(left, right chan int) {
	r := <-right
	l := 1 + r
	fmt.Printf("(%d) <- (%d + 1)\n", l, r)
	left <- l
}

func main() {
	const n = 1e1
	leftmost := make(chan int)
	var right chan int
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1 // seed
	}(right)
	fmt.Printf("finally: %d\n", <-leftmost)
}
