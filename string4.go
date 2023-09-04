package main

import "fmt"

func main() {

	var numeros []float32
	var numero float32
	var soma float32

	for i := 0; i < 6; i++ {

		fmt.Print("Digite um número: ")
		fmt.Scan(&numero)

		numeros = append(numeros, numero)
	}
	fmt.Println(numeros)

	for _, num := range numeros {
		soma += num
	}
	fmt.Print(soma / float32(len(numeros)))
}
