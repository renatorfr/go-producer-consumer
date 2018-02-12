package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	maxProducers := 5
	var maxConsumers = 6

	Run(maxProducers, maxConsumers)
}

func Run(maxProducers int, maxConsumers int) {
	var wg sync.WaitGroup

	delivery := make(chan food)

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

func producer(id int, delivery chan<- food, wg *sync.WaitGroup) {
	defer wg.Done()

	// fmt.Printf("Producer created: %v\n", id)

	var f food

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(2) == 1 {
		f = nutella{weight: id}
		// fmt.Printf("Produced nutella: %+v\n", f)
	} else {
		f = tomato{weight: id}
		// fmt.Printf("Produced tomato: %+v\n", f)
	}

	select {
	case delivery <- f:
	case <-time.After(2 * time.Second):
		// fmt.Printf("Timeout producer: %v\n", id)
	}
}

func consumer(id int, delivery <-chan food, wg *sync.WaitGroup) {
	defer wg.Done()

	// fmt.Printf("Consumer created: %v\n", id)

	select {
	case n := <-delivery:
		n.eat()
	case <-time.After(2 * time.Second):
		// fmt.Printf("Timeout consumer: %v\n", id)
	}
}

type nutella struct {
	weight int
}

func (n nutella) eat() {
	// fmt.Printf("Eating %v grams of nutella\n", n.weight)
}

type food interface {
	eat()
}

type tomato struct {
	weight int
}

func (r tomato) eat() {
	// fmt.Printf("Eating %v grams of tomato\n", r.weight)
}
