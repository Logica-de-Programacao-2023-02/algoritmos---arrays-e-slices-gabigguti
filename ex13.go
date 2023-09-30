package main

import "fmt"

func main() {
	ele := [7]int{2, 3, 5, 9, 12, 14, 17}
	var n1, n2 int
	fmt.Println(ele)
	fmt.Println("Informe o número que será adicionado no primeiro: ")
	fmt.Scanln(&n1)
	fmt.Println("Informe o número que será adicionado no ultimo: ")
	fmt.Scanln(&n2)
	l1 := make([]int, 0, len(ele)+2)
	l1 = append(l1, n1)
	l1 = append(l1, ele[:]...)
	l1 = append(l1, n2)
	fmt.Println(l1)
}
