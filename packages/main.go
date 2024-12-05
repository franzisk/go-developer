package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func example1() {
	fmt.Println("\n========== 1 ")
	fmt.Println(strings.Contains("Maria da Silva", "Silva"))
	fmt.Println(strings.Contains("Maria da Silva", "Pereira"))
}

func example2() {
	fmt.Println("\n========== 2 ")
	if _, err := io.WriteString(os.Stdout, "Hello World\n"); err != nil {
		log.Fatal(err)
	}
}

func example3() {
	fmt.Println("\n========== 3 ")
	message := []byte("Hello, Gophers!")
	err := os.WriteFile("hello.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func example4() {
	fmt.Println("\n========== 4 ")
	texto := "exemplo de título com unicode: café & música"

	// Cria um "caser" para formatação de título
	titleCaser := cases.Title(language.Und) // "Und" é para texto sem idioma específico

	// Converte o texto para o formato de título
	resultado := titleCaser.String(texto)

	fmt.Println("Texto original:", texto)
	fmt.Println("Texto formatado:", resultado)
}

func example5() {
	fmt.Println("\n========== 5 ")
	file, error := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0644)
	if error != nil {
		log.Fatal(error)
	}
	if error := file.Close(); error != nil {
		log.Fatal(error)
	}
}

func example6() {

}

func main() {
	fmt.Println("====== PACOTES =========")
	example1()
	example2()
	example3()
	example4()
	example5()
}
