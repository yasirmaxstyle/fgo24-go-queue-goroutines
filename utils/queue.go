package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func queue(wg *sync.WaitGroup, done chan float64) {
	defer wg.Done()
	start := time.Now()
	delay := rand.Intn(5-1+1) + 1
	time.Sleep((time.Second) * time.Duration(delay))
	end := time.Now()
	duration := end.Sub(start) // time ellapses
	done <- float64(duration)
}

func Queue() {
	total := rand.Intn(10-1+1) + 1

	wg := sync.WaitGroup{}     // make waitgroup
	done := make(chan float64) // make chanel

	fmt.Printf("--- You have %d queues ---\n", total)

	for x := range total {
		wg.Add(1)
		go queue(&wg, done)
		fmt.Printf("Queue %d: ", x+1)
		time := float64(<-done) / 1e9 //divided by 1 exp 9 to get rid of "e+" in result
		fmt.Printf("%.9f\n", time)
		wg.Wait()
	}

	fmt.Println("All queues completed!")
}
