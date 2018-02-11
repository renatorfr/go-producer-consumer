package main

import (
	"fmt"
	"sync"
)

func main() {
	maxProducers := 5
	var maxConsumers = 5

	var wg sync.WaitGroup

	delivery := make(chan string)

	for i := 1; i <= maxProducers; i++ {
		wg.Add(1)
		go producer(i, delivery, &wg)
	}

	for i := 1; i <= maxConsumers; i++ {
		wg.Add(1)
		go consumer(i, delivery, &wg)
	}

	wg.Wait()
}

func producer(id int, delivery chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Producer created: %v\n", id)

	p := fmt.Sprintf("Product %d\n", id)
	fmt.Printf("Produced: %v", p)
	delivery <- p
}

func consumer(id int, delivery <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Consumer created: %v\n", id)

	p := <-delivery
	fmt.Printf("Consumed: %v", p)
}
