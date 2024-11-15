package main

import "fmt"

const ebulicaoK = 273.15

func main() {
	var temperaturaK = 300.0
	var tempC = temperaturaK - ebulicaoK
	fmt.Printf("°K %.2f em °C é %.2f\n", temperaturaK, tempC)
}
