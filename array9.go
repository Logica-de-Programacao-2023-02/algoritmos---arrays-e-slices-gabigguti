// Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em todas as posições do Array. Imprima o Array resultante

package main

import "fmt"

func main() {
	var array = [6]float64{1.4, 1.6, 1.8, 2.2, 2.4, 2.6}

	var x float64
	fmt.Print("Digite um numero:")
	fmt.Scan(&x)

	for i := 0; i < len(array); i++ {
		array[i] = array[i] + x
	}
	fmt.Println(array)

}
