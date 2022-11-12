package main

import (
	"fmt"
	"sync"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
	// mutex   sync.Mutex     // mutex is used to define a critical section of code.
)

func main() {
	// wg.Add(3) // Add a count of two, one for each goroutine.
	ch := make(chan string)
	defer close(ch)
	arr := []string{"Python", "GO", "Java"}
	// go python(ch, "Python")

	// go python(ch, "Go Programming Language")
	// go python(ch, "Java")
	// fmt.Println("Counter:", counter)
	for _, item := range arr {
		wg.Add(2)
		go python(ch, item)
		go readLang(ch)
	}
	fmt.Println("Read channel : ", <-ch)
	_, ok := <-ch
	if ok {
		fmt.Println("Channel Opened")
		// defer close(ch)
	} else {
		fmt.Println("Channel closed")
	}

	// time.Sleep(time.Second * 2)
	wg.Wait() // Wait for the goroutines to finish.
	// close(ch)
	// for item := range ch {
	// 	fmt.Println(item)
	// }

}

func python(ch chan<- string, lang string) {
	// fmt.Println(lang)
	defer wg.Done() // Schedule the call to Done to tell main we are done.
	ch <- lang
	// defer close(ch)
	// for i := 0; i < 3; i++ {
	// 	mutex.Lock()
	// 	{
	// 		// fmt.Println(lang)
	// 		ch <- lang
	// 		counter++
	// 	}
	// 	mutex.Unlock()

	// }
}

func program(ch chan<- string, lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.
	ch <- lang
	// defer close(ch)
}

func java(ch chan<- string, lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.
	ch <- lang
	// defer close(ch)
}

func readLang(ch <-chan string) {
	defer wg.Done()
	fmt.Println(<-ch)

}
