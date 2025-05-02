package main

import ( 
	"fmt"
    "math/rand"
	"sync"
    "time"
)

func main() {
	inCh := make(chan int, 5)
    outCh := make(chan int, 5)
	rand.Seed(time.Now().UnixNano())
	var wgProducers sync.WaitGroup
	var wgConsumers sync.WaitGroup
	var wgFinal sync.WaitGroup

	wgProducers.Add(2)

	// Two Producers
	go func() {
		defer wgProducers.Done()
		for i := 1; i <= 29; i += 2 {
			inCh <- i
			time.Sleep(time.Duration(rand.Intn(1501)) * time.Millisecond)
		}
	}()

	go func() {
		defer wgProducers.Done()
		for i := 2; i <= 30; i += 2 {
			inCh <- i
			time.Sleep(time.Duration(rand.Intn(1501)) * time.Millisecond)
		}
	}()

	go func() {
		wgProducers.Wait()
		close(inCh)
	}()
	
	wgConsumers.Add(2)

	// Two Consumers
	go func() {
		defer wgConsumers.Done()
		for value := range inCh {
			outCh <- (value * value)
			time.Sleep(time.Duration(rand.Intn(3001)) * time.Millisecond)
		}
	}()

	go func() {
		defer wgConsumers.Done()
		for value := range inCh {
			outCh <- (value * value)
			time.Sleep(time.Duration(rand.Intn(3001)) * time.Millisecond)
		}
	}()

	go func() {
		wgConsumers.Wait()
		close(outCh)
	}()
	
	wgFinal.Add(1)

	// Final Filter
	go func() {
		defer wgFinal.Done()
		previous := -1
		for value := range outCh {
			if value > previous {
				fmt.Println(value)
				previous = value
			}
		}
	}()

	wgFinal.Wait()
}