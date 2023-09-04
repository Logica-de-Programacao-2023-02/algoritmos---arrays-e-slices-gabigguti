package main

import "fmt"

func main() {
	var num = [5]int{1, 2, 3, 4, 5}
	for _, numero := range num {
		fmt.Println(numero)
	}
}
