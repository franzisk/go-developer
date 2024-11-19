package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// Calcular a área de um retângulo qualquer
// Recebe um ponteiro
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

// Calcular o perímetro de um retângulo qualquer
// Recebe um valor
func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func retanguloImp() {
	r := retangulo{comprimento: 10, altura: 5}
	fmt.Println("Área:", r.area())
	fmt.Println("Perímetro:", r.perimetro())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimetro())
}

func main() {
	retanguloImp()
}
