package main

import (
	"fmt"
)

func arrays() {
	fmt.Println("========================")
	fmt.Println("ARRAYS")
	var grades [5]float64
	grades[0] = 5.3
	grades[1] = 8.0
	grades[2] = 6.8
	grades[3] = 9.9
	grades[4] = 7.5
	var total float64 = 0.0
	for i := 0; i < len(grades); i++ {
		total = total + grades[i]
		fmt.Println("nota", grades[i])
	}
	fmt.Println("média", total/float64(len(grades)))
}

func slices() {
	fmt.Println("========================")
	fmt.Println("SLICES")
	arr := [5]float64{10, 20, 30, 40, 50}

	slice1 := arr[0:5]
	fmt.Println("fatia 1", slice1)

	slice2 := append(slice1, 50.33)
	fmt.Println("fatia 2", slice2)

	slice3 := make([]float64, 2)
	fmt.Println("fatia 3 (1)", slice3)

	copy(slice3, slice1)
	fmt.Println("fatia 3 (2)", slice3)
}

func maps() {
	fmt.Println("========================")
	fmt.Println("MAPS")
	m := make(map[string]float64)

	m["k1"] = 7.567
	m["k2"] = 13.754

	fmt.Println("map:", m)
	fmt.Println("chave 1:", m["k1"])
	fmt.Println("chave 2:", m["k2"])

	x := make([]int, 3, 9)
	fmt.Println(x)

}

func main() {
	arrays()
	slices()
	maps()
}
