package main

import "fmt"

func main() {
	slice := []int{1, 3, 6, 9, 12, 15, 19, 20}
	fmt.Println(slice)
	var x, y int
	fmt.Println("Digite o elemento ")
	fmt.Scan(&x)
	x -= 1
	fmt.Println("Digite o segundo elemnto")
	fmt.Scan(&y)
	y -= 1
	slice[x], slice[y] = slice[y], slice[x]
	fmt.Println(slice)
}
