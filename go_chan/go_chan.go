package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Go routine")

	fmt.Println("Main")
	/*
		for i := 0; i < 3; i++ {
			go func(n int) {
				fmt.Println(n)
			}(i)
		}

		for j := 0; j < 3; j++ {
			j := j
			go func() {
				fmt.Println(j)
			}()
		}

		ch := make(chan string)

		go func() {
			ch <- "Message to main channel from an anonymous function"
		}()
		msg := <-ch
		fmt.Println(msg)
	*/
	/*
		go func() {
			for i := 0; i < 3; i++ {
				ch <- fmt.Sprintf("Message %d", i+1)
			}
			close(ch)
		}()

		for msg := range ch {
			fmt.Println("got message", msg)
		}

		message, ok := <-ch
		fmt.Printf("Was there a real message? %#v %v \n", message, ok)
	*/
	// For every value n in values spin a Goroutine that will
	//sleep for n millisecondsc
	//Send n over a channel
	//In the function body collect values from a channel to a slice and return it

	values := []int{15, 8, 9, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))

	time.Sleep(10 * time.Millisecond)
}

func sleepSort(values []int) []int {
	ch := make(chan int)
	var messages []int
	for i := range values {
		go func(n int) {
			time.Sleep(time.Duration(values[n]) * time.Millisecond)
			ch <- values[n]
		}(i)
		//msg := <-ch
		//messages = append(messages, msg)
	}

	for range values {
		n := <-ch
		messages = append(messages, n)

	}

	return messages
}
