package main

import "fmt"

func fazerCoisa(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go fazerCoisa(10)

	var escrever string

	fmt.Scanln(&escrever)
}
