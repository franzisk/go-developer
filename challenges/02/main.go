package main

import "fmt"

const maximo int = 100

func parte1() {
	fmt.Println("====== DESAFIO #2 - PARTE 1")
	for i := 0; i < maximo; i++ {
		if i%3 == 0 {
			fmt.Println(i, "é divisível por 3")
		}
	}
}

func parte2() {
	fmt.Println("====== DESAFIO #2 - PARTE 2")
	for i := 0; i < maximo; i++ {
		var saida string

		if i%3 == 0 {
			saida += "Pin"
		}
		if i%5 == 0 {
			saida += "Pan"
		}

		if saida != "" {
			fmt.Println(saida) // Imprime "Pin", "Pan" ou "PinPan"
		} else {
			fmt.Println(i) // Imprime o número se não for múltiplo de 3 ou 5
		}
	}
}

func main() {
	parte1()
	parte2()
}
