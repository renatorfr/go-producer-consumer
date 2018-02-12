package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	maxProducers := 5
	var maxConsumers = 5

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

	fmt.Printf("Producer created: %v\n", id)

	if rand.Intn(3) > 1 {
		var n nutella
		n.weight = id
		fmt.Printf("Produced nutella: %+v\n", n)

		delivery <- n
	} else {
		var t tomato
		t.weight = id
		fmt.Printf("Produced tomato: %+v\n", t)

		delivery <- t
	}

}

func consumer(id int, delivery <-chan food, wg *sync.WaitGroup) {
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

type food interface {
	eat()
}

type tomato struct {
	weight int
}

func (r tomato) eat() {
	fmt.Printf("Eating %v grams of tomato\n", r.weight)
}
