package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

	var f food

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(2) == 1 {
		f = nutella{weight: id}
		fmt.Printf("Produced nutella: %+v\n", f)
	} else {
		f = root{weight: id}
		fmt.Printf("Produced root: %+v\n", f)
	}

	delivery <- f
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

type root struct {
	weight int
}

func (r root) eat() {
	fmt.Printf("Eating %v grams of root\n", r.weight)
}
