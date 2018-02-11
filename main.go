package main

import (
	"fmt"
	"sync"
)

func main() {
	maxProducers := 5
	var maxConsumers = 5

	var wg sync.WaitGroup

	delivery := make(chan nutella)

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

func producer(id int, delivery chan<- nutella, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Producer created: %v\n", id)

	var n nutella
	n.weight = id
	fmt.Printf("Produced: %+v\n", n)

	delivery <- n
}

func consumer(id int, delivery <-chan nutella, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Consumer created: %v\n", id)

	n := <-delivery
	n.eat()
}

type nutella struct {
	weight int
}

func (n nutella) eat() {
	fmt.Printf("Eating %v grams of nutella\n", n.weight)
}
