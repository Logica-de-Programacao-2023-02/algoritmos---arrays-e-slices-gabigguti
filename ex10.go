package main

import "fmt"

func main() {
	num := []int{4, 3, 5, 7, 8, 9, 11, 22, 17, 21}
	fmt.Println(num)
	min := num[0]
	max := num[0]
	for _, eles := range num {
		if eles < min {
			min = eles
		}
		if eles > max {
			max = eles
		}
	}
	fmt.Println("O menor número é: ", min)
	fmt.Println("O maior número é: ", max)

}
