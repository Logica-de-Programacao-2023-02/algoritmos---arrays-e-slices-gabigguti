package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite o valor de n: ")
	fmt.Scan(&n)
	var primos []int
	for i := n; i >= 1; i-- {
		if n%i == 0 && n%n == 0 {
			primos = append(primos, i)
		} else if n%i != 0 || n%n != 0 {
			continue
		}
	}
	fmt.Println(primos)
}
