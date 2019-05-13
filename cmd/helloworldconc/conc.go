package main

import (
	"fmt"
)

//goroutine 例子
func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		//go start a goroutine
		go printHelloWorld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("hello world from goroutine %d\n", i)
	}
}
