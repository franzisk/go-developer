package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	// Obtém o horário atual
	now := time.Now()

	// Formata para o formato desejado
	// 02: Representa o dia (DD).
	// 01: Representa o mês (MM).
	// 2006: Representa o ano com 4 dígitos (YYYY).
	// 15: Representa a hora em formato de 24 horas (HH).
	// 04: Representa os minutos (MM).
	// 05: Representa os segundos (SS).
	formattedTime := now.Format("02/01/2006 15:04:05")

	return fmt.Sprintf("ERROR: %v: %v", formattedTime, e.What)
}

func erroTest() error {
	return MyError{
		time.Now(),
		"the file system has gone away",
	}
}

func main() {
	if err := erroTest(); err != nil {
		fmt.Println(err)
	}

	err := MyError{
		When: time.Now(),
		What: "Failed to connect to the database",
	}
	fmt.Println(err) // ERROR: 2024-11-25 14:15:00 +0000 UTC: Failed to connect to the database
}
