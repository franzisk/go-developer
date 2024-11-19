package main

import "fmt"

type person struct {
	name string
	age  int
}

type retangulo struct {
	comprimento, altura int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 33
	return &p
}

func personImpl() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
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
	personImpl()
	retanguloImp()
}
