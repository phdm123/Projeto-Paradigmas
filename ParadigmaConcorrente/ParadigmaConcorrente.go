package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Cria as duas goroutines, passando o valor do "id" como argumento
	go doWork(1)
	go doWork(2)

	// Aguarda pelo término das goroutines
	time.Sleep(2 * time.Second)
}
