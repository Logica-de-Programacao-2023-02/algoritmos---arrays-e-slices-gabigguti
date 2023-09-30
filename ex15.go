package main

import "fmt"

func main() {
	var ele = [10]float64{1.1, 2.3, 5.5, 4.5, 7.6, 8.6, 10.5, 11.2, 13.5, 9.5}
	fmt.Println(ele)
	var result = []float64{}
	for _, eles := range ele {
		if eles > 5 {
			result = append(result, eles)
		}
	}
	fmt.Println(result)
}
