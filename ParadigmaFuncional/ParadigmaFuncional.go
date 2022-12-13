package main

import "fmt"

func main() {
	// Declara uma função anônima que recebe um número inteiro e retorna o seu quadrado
	square := func(x int) int {
		return x * x
	}

	// Utiliza a função anônima para calcular o quadrado de cada elemento de um slice de números inteiros
	numbers := []int{1, 2, 3, 4, 5}
	squaredNumbers := make([]int, len(numbers))
	for i, n := range numbers {
		squaredNumbers[i] = square(n)
	}

	// Imprime o resultado
	fmt.Println(squaredNumbers) // [1, 4, 9, 16, 25]
}
