package main

import (
	"fmt"
	"time"
)

//function to print hello
func printHello(ch chan int) {
	fmt.Println("Hello from printHello")
	ch <- 2
}

func main() {
	fmt.Println("Time in s : ", time.Second)
	end_time := time.Now().Unix()
	fmt.Println(end_time)
	//make a channel. You need to use the make function to create channels.
	//channels can also be buffered where you can specify size. eg: ch := make(chan int, 2)
	ch := make(chan int)
	//inline goroutine. Define a function inline and then call it.
	go func() {
		fmt.Println("Hello inline")
		ch <- 1
	}()
	//call a function as goroutine
	go printHello(ch)
	fmt.Println("Hello from main")

	//get first value from channel.
	//and assign to a variable to use this value later
	//here that is to print it

	i := <-ch
	fmt.Println("Received ", i)

	//get the second value from channel
	//do not assign it to a variable because we dont want to use that

	<-ch
}
