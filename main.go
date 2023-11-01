package main

import (
	"fmt"
	"sync"
	"cloud-bot/resources"
)

func transferData() {
	// Create a channel for communication
	ch := make(chan int)

	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Goroutine 1: Send data to the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i // Send data to the channel
		}
		close(ch) // Close the channel to signal that no more data will be sent
	}()

	// Goroutine 2: Receive data from the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("Received: %d\n", num)
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}

func main() {
	fmt.Println("Starting trading bot")
	resources.login()
}
