package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4, 5}
	num = append(num[:2], num[3:]...)
	fmt.Println(num)
}
