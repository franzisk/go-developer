package main

import (
	"fmt"
	"time"
)

func ping(pingChannel chan string, quitChannel chan bool) {
	for {
		select {
		case pingChannel <- "ping":
			time.Sleep(time.Second) // Pequeno atraso para regular o envio
		case <-quitChannel: // Recebe sinal de encerramento
			close(pingChannel) // Fecha o canal de `ping`
			return             // Sai da goroutine
		}
	}
}

func pong(pongChannel chan string, quitChannel chan bool) {
	for {
		select {
		case pongChannel <- "pong":
			time.Sleep(time.Second) // Pequeno atraso para regular o envio
		case <-quitChannel: // Recebe sinal de encerramento
			close(pongChannel) // Fecha o canal de `pong`
			return             // Sai da goroutine
		}
	}
}

// Fazer a impressão do ping e pong alternando entre os canais
func printMessages(pingChannel chan string, pongChannel chan string, quitChannel chan bool) {
	for {
		select {
		case message := <-pingChannel:
			fmt.Println(message)
		case message := <-pongChannel:
			fmt.Println(message)
		case <-quitChannel:
			fmt.Println("Encerrando...")
			return
		}
	}
}

func main() {
	pingChannel := make(chan string)
	pongChannel := make(chan string)

    // Um canal adicional (quitChannel) permite sinalizar a todas as goroutines para encerrarem.
	quitChannel := make(chan bool)

	go ping(pingChannel, quitChannel)
	go pong(pongChannel, quitChannel)
	go printMessages(pingChannel, pongChannel, quitChannel)

	// Aguardar entrada do usuário para encerrar
	fmt.Println("Pressione ENTER para encerrar...")
	fmt.Scanln()

	// Sinaliza para as goroutines encerraram, depois de uma tecla pressionada
	close(quitChannel)
	time.Sleep(time.Second) // Tempo para as goroutines terminarem
}
