package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	names := []string{
		"Leidiane",
		"André Leidiane", 
		"Helena", 
		"Beatriz", 
		"André Creonice",
		"Mayara", 
		"Tadeu", 
		"Sophia", 
		"Creonice", 
		"Varte", 
		"Luciano", 
		"Serma",
		"Vinícius", 
		"Camili", 
		"Vitória", "Nadir", "Sérgio", "Maurício",
		"Francisco", "Sônia", "Michele",
	}

	// Garantir aleatoriedade
	rand.Seed(time.Now().UnixNano())

	// Criar uma cópia da lista para o sorteio
	drawPool := make([]string, len(names))
	copy(drawPool, names)

	// Resultado do sorteio
	results := make(map[string]string)

	// Realizar o sorteio
	for _, name := range names {
		possibleMatches := filterOut(drawPool, name)
		if len(possibleMatches) == 0 {
			fmt.Println("Erro: Não foi possível completar o sorteio sem repetições. Tente novamente.")
			return
		}

		// Escolher um nome aleatório dos possíveis
		drawIndex := rand.Intn(len(possibleMatches))
		drawnName := possibleMatches[drawIndex]

		// Registrar o sorteio
		results[name] = drawnName

		// Remover o sorteado do pool
		drawPool = remove(drawPool, drawnName)
	}

	// Exibir os pares do sorteio com destaque
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("🔥 Resultado do Amigo-Secreto 🔥")
	fmt.Println(strings.Repeat("=", 30))
	for giver, receiver := range results {
		fmt.Printf("🎁 %-20s ➡️    %-20s 🎉\n", giver, receiver)
	}
	fmt.Println(strings.Repeat("=", 30))
}

// Filtra a lista removendo o nome do sorteador
func filterOut(pool []string, name string) []string {
	filtered := []string{}
	for _, n := range pool {
		if n != name {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

// Remove um nome específico da lista
func remove(pool []string, name string) []string {
	for i, n := range pool {
		if n == name {
			return append(pool[:i], pool[i+1:]...)
		}
	}
	return pool
}
