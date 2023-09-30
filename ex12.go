package main

import "fmt"

func main() {
	num := [5]int{7, 5, 6, 17, 21}
	nv := []int{}
	for _, eles := range num {
		if eles%3 == 0 {
			nv = append(nv, eles)
		}
	}
	fmt.Println(num)
	fmt.Println(nv)
}
