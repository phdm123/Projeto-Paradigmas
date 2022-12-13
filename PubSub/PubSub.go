package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var pipe = make(chan int, 10)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go produtor()
	go consumidor()
	go consumidor()
	wg.Wait()
}

func produtor() {
	for {
		//var item int = 1
		item := 1
		time.Sleep((time.Second) * time.Duration(rand.Intn(3)))
		fmt.Printf("PRODUTOR: adicionei um item! Pipe length: %v\n", len(pipe))
		pipe <- item
	}
}

func consumidor() {
	for {
		time.Sleep((time.Second) * time.Duration(rand.Intn(7)))
		item := <-pipe
		fmt.Printf("Consumidor: consumi um item %v , pipe length: %v \n", item, len(pipe))
	}
}
