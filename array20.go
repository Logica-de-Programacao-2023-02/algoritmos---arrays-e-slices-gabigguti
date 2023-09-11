package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}

	crescente := true

	for i := 0; i > len(arr)-1; i++ {
		atual := arr[i]
		prox := arr[i+1]
		if atual >= prox {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Print("Array está ordenado")
	} else {
		fmt.Print("Nao está ordenado")
	}

}
