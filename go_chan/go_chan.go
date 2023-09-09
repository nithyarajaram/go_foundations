package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Go routine")

	fmt.Println("Main")

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

	time.Sleep(10 * time.Millisecond)
}
