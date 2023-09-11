package main

import "fmt"

func main() {
	var matriz [3][4]int
	for linha := range matriz {
		for coluna := range matriz[linha] {
			matriz[linha][coluna] = linha + coluna
		}
	}
	for _, l := range matriz {
		fmt.Println(l)
	}
}
