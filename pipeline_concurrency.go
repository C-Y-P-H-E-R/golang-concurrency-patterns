package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Stage 1: Generate random numbers.
func generate(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case out <- rand.Intn(100):
				// Send a random number.
			case <-done:
				return
			}
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate work.
		}
	}()
	return out
}

// Stage 2: Square the numbers.
func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// Stage 3: Filter numbers greater than 50 (squared values > 50).
func filter(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				return
			default:
				if n > 50 {
					select {
					case out <- n:
					case <-done:
						return
					}
				}
			}
		}
	}()
	return out
}

// Stage 4: Print the numbers.
func printer(done <-chan struct{}, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		select {
		case <-done:
			return
		default:
			fmt.Println(n)
		}
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	
	// Create the pipeline
	nums := generate(done)
	squared := square(done, nums)
	filtered := filter(done, squared)
	
	// Start printer
	go printer(done, filtered, &wg)
	
	// Let it run for a short time
	time.Sleep(3 * time.Second)
	
	// Signal stages to stop
	close(done)
	
	// Wait for printer to finish
	wg.Wait()
	fmt.Println("Pipeline completed.")
}
