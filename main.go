package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	maxProducers := 1
	var maxConsumers = 10

	delivery := make(chan food)

	var wg sync.WaitGroup
	wg.Add(maxProducers)
	wg.Add(maxConsumers)

	for i := 1; i <= maxProducers; i++ {
		go producer(i, delivery, &wg)
	}

	for i := 1; i <= maxConsumers; i++ {
		go consumer(i, delivery, &wg)
	}

	wg.Wait()
}

type food interface {
	eat()
}

type nutella struct {
	Size int
}

func (n nutella) eat() {
	fmt.Printf("Eating Nutella!!! %v\n", n.Size)
	time.Sleep(time.Second * time.Duration(n.Size))
}

func producer(id int, delivery chan<- food, wg *sync.WaitGroup) {
	fmt.Printf("Producer created: %v\n", id)

	var n nutella
	n.Size = rand.Intn(20)

	delivery <- n

	wg.Done()
}

func consumer(id int, delivery <-chan food, wg *sync.WaitGroup) {
	fmt.Printf("Consumer created: %v\n", id)

	n := <-delivery

	n.eat()

	wg.Done()
}
