package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxProducers := 5
	var maxConsumers = 5

	delivery := make(chan string)

	for i := 1; i <= maxProducers; i++ {
		go producer(i, delivery)
	}

	for i := 1; i <= maxConsumers; i++ {
		go consumer(i, delivery)
	}
}

func producer(id int, delivery chan<- string) {
	fmt.Printf("Producer created: %v\n", id)

	delivery <- fmt.Sprintf("Product %d", rand.Intn(20))
}

func consumer(id int, delivery <-chan string) {
	fmt.Printf("Consumer created: %v\n", id)

	p := <-delivery

	fmt.Printf("Product consumed: %s", p)
}
