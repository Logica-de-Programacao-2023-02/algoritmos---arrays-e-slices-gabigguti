package main

import "fmt"

func main() {
	ele := [10]int{1, 1, 8, 12, 15, 23, 33, 51, 45, 13}
	fmt.Println(ele)
	soma := 0
	for i := 1; i < len(ele); i += 2 {
		soma += ele[i]
	}
	fmt.Println(soma)
}
