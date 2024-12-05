package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Função para simular o download de um arquivo
func downloadFile(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca a tarefa como concluída no WaitGroup

	fmt.Printf("Iniciando download do arquivo %d...\n", id)
	// Simula um tempo aleatório de download
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	fmt.Printf("Download do arquivo %d concluído!\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Gera números aleatórios para simular tempos diferentes

	var wg sync.WaitGroup // Cria um WaitGroup para gerenciar as goroutines

	numDownloads := 5 // Número de downloads simulados

	// Inicia múltiplas goroutines para simular downloads simultâneos
	for i := 1; i <= numDownloads; i++ {
		wg.Add(1) // Incrementa o contador do WaitGroup
		go downloadFile(i, &wg)
	}

	fmt.Println("Todos os downloads foram iniciados...")
	wg.Wait() // Aguarda todas as goroutines terminarem
	fmt.Println("Todos os downloads foram concluídos!")
}
