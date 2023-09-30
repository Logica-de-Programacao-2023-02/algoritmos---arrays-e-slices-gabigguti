package main

import "fmt"

func main() {
	n := [4]float64{1.4, 8.4, 4.7, 9.4}
	var prod float64 = 1
	for i := 0; i < len(n); i++ {
		prod *= n[i]
	}
	fmt.Println(prod)
}
