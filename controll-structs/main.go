package main

import "fmt"

func instfor() {
	fmt.Println("========================")
	fmt.Println("Instrução FOR")
	for i := 0; i < 10; i++ {
		fmt.Println(`O valor de i = `, i)
	}
}

func intsif() {
	fmt.Println("========================")
	fmt.Println("Instrução IF")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, `é par `)
		} else {
			fmt.Println(i, `é ímpar `)
		}
	}
}

func instswitch() {
	fmt.Println("========================")
	fmt.Println("Instrução SWITCH")
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "é par")
		default:
			fmt.Println(i, "é ímpar")
		}
	}
}

func booleanops() {
	fmt.Println("========================")
	fmt.Println("Operador BOOLEANO")
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("Verdade 1")
	}
	if x != 2 || x == 3 {
		fmt.Println("Verdade 2")
	}
}

func looping1() {
	fmt.Println("========================")
	fmt.Println("LOOPINGS (1)")
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func looping2() {
	fmt.Println("========================")
	fmt.Println("LOOPINGS (2)")
	for horas := 0; horas <= 12; horas++ {
		fmt.Println(horas, "hora")
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Println(horas, ":", minutos)
		}
	}
}

func looping3() {
	fmt.Println("========================")
	fmt.Println("LOOPINGS (3) (while)")
	i := 0
	for i <= 10 {
		fmt.Println(i, "dentro do valor")
		i++
	}
}

func looping4() {
	fmt.Println("========================")
	fmt.Println("LOOPINGS (4) (infinito)")
	i := 0
	for {
		if i < 20 {
			fmt.Println(i, "dentro do valor")
			i++
			continue
		} else {
			break
		}
	}
}

func main() {
	instfor()
	intsif()
	instswitch()
	booleanops()
	looping1()
	looping2()
	looping3()
	looping4()
}
