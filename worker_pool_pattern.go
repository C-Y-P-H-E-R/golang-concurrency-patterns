package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, taskChan chan int) {
	defer wg.Done()
	for i := range taskChan {
		t := time.Now()
		fmt.Println(`Task `, i, ` Processing`)
		// simulate work
		time.Sleep(time.Duration(200) * time.Millisecond)
		elapsed := time.Since(t)
		fmt.Println(`Task `, i, ` Finished. Time elapsed = `, elapsed)
	}
}

func main() {
	tasks := 200
	workers := 10
	// the tasks must come via a queue (confinement)
	// unbuffered channel to simulate real world channel
	taskChan := make(chan int)
	var wg sync.WaitGroup

	// shooting workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		// need to pass task channel so that worker can keep tabs on how many tasks left
		go Worker(&wg, taskChan)
	}

	for i := 0; i < tasks; i++ {
		taskChan <- i
	}

	close(taskChan)
	wg.Wait()
}
