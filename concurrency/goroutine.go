package concurrency

import (
	"fmt"
	"time"
)

// NOTE: Goroutines run in the same address space, so access to shared memory must be synchronized.

// NOTE: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

// NOTE: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

func Goroutine() {
	go say("Hi")
	say("Hello")

	channel()
	bufferedChannel() // Channel with multiple values
	SelectChannel()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channel() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Channel variable can receive multiple values
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fibChannel := make(chan int, 10)
	go fibonacci(cap(fibChannel), fibChannel)
	for i := range fibChannel {
		fmt.Println(i)
	}
}

func bufferedChannel() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/*
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/
func SelectChannel() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
}

func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		// NOTE: Use a default case to try a send or receive without blocking
		default:
			// receiving from c would block
		}
	}
}
