package main

import "fmt"

func main() {
	ele := [10]int{1, 5, 8, 12, 14, 20, 34, 50, 45, 11}
	fmt.Println(ele)
	result := []int{}
	for _, eles := range ele {
		if eles%2 == 0 {
			result = append(result, eles)
		}
	}
	fmt.Println(result)
}
