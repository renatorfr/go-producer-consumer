package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var nutellasEaten = 0
var tomatoesEaten = 0
var mutex = &sync.Mutex{}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxProducers := rand.Intn(500) + 1
	var maxConsumers = rand.Intn(500) + 1

	fmt.Printf("Max producers: %v | Max consumers: %v\n", maxProducers, maxConsumers)

	Run(maxProducers, maxConsumers)

	fmt.Printf("Nutellas eaten: %v | Tomatoes eaten: %v\n", nutellasEaten, tomatoesEaten)
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

	var f food

	rand.Seed(time.Now().UnixNano())

	if rand.Intn(2) == 1 {
		f = nutella{weight: id}
	} else {
		f = tomato{weight: id}
	}

	select {
	case delivery <- f:
	case <-time.After(2 * time.Second):
	}
}

func consumer(id int, delivery <-chan food, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case n := <-delivery:
		n.eat()
	case <-time.After(2 * time.Second):
	}
}

type nutella struct {
	weight int
}

func (n nutella) eat() {
	mutex.Lock()
	nutellasEaten++
	mutex.Unlock()
}

type food interface {
	eat()
}

type tomato struct {
	weight int
}

func (r tomato) eat() {
	mutex.Lock()
	tomatoesEaten++
	mutex.Unlock()
}
