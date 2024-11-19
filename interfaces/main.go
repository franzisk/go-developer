package main

import (
	"fmt"
	"math"
)

// coleção de métodos da interface
// somente depois eles vão ser declarados
type geometria interface {
	area() float64
	perimetro() float64
}

// Aplicar a interface geometria nos dois tipos abaixo
type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

// Ao declarar o método "area" que está na interface "geometria"
// estamos automaticamente usando a interface "geometria"
// Então aqui estamos usando a interface "geometria" no tipo "quadrado"
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// ==============================================
// Então aqui estamos usando a interface "geometria" no tipo "circulo"
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Println("======== MEDIR =====")
	fmt.Println("geometria:", g)
	fmt.Println("área", g.area())
	fmt.Println("perímetro", g.perimetro())
}

func main() {

	// o tipo "quadrado" implementa a interface "geografia", então pode ser passado para a função
	quadrado1 := quadrado{lado: 20}
	medir(quadrado1)

	// o tipo "circulo" implementa a interface "geografia", então pode ser passado para a função
	circulo1 := circulo{raio: 15}
	medir(circulo1)
}
