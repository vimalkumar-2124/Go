// a mutex, or a mutual exclusion is a mechanism that allows us to prevent concurrent processes from entering a critical section of data whilst it’s already being executed by a given process.

package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000
}

func withDraw(value int, wg *sync.WaitGroup) {
	// time.Sleep(5 * time.Second)
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance : %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func deposit(value int, wg *sync.WaitGroup) {
	// time.Sleep(2 * time.Second)
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance : %d \n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}
func main() {
	fmt.Println("GO mutex started")
	var wg sync.WaitGroup
	wg.Add(2)
	go withDraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("New balance %d \n", balance)
}
