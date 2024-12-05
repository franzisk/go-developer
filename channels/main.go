package main

import (
	"fmt"
	"time"
)

// Parâmetro:
// pingChannel chan string é um canal usado para transmitir dados do tipo string.
func ping(pingChannel chan string) {
	for i := 0; ; i++ { // Loop infinito
		// insere a string "ping" no canal pingChannel.
		// A operação bloqueia até que alguém consuma a mensagem do canal.
		pingChannel <- "ping" // Envia a string "ping" ao canal
	}
}

func print(pingChannel chan string) {
	for { // Loop infinito: Executa continuamente enquanto o programa estiver rodando.
		// Lê uma mensagem do canal pingChannel.
		// A operação bloqueia até que algo esteja disponível no canal.
		message := <-pingChannel // Recebe uma string do canal

		fmt.Println(message) // Imprime a mensagem no console

		time.Sleep(time.Second * 1) // Aguarda 1 segundo antes de repetir
	}
}

func main() {
	// Cria um canal não-bufferizado para transmitir strings entre goroutines.
	var channel chan string = make(chan string) // Cria um canal para strings

	// Inicia a execução da função ping em uma goroutine separada.
	go ping(channel) // Inicia a goroutine `ping`

	// Inicia a execução da função print em uma goroutine separada.
	go print(channel) // Inicia a goroutine `print`

	// Estas goroutines compartilham o canal channel para comunicação.

	var inputText string

	// Bloqueia a execução principal até que o usuário insira algo e pressione Enter.
	// Isso mantém o programa rodando para que as goroutines possam continuar funcionando.
	fmt.Scanln(&inputText) // Aguarda entrada do usuário para encerrar o programa

	// NOTA:
	// O fmt.Scanln é usado no código original para impedir que o
	// programa principal termine imediatamente. Sem ele, o programa
	// termina antes que as goroutines tenham a chance de produzir qualquer saída.
}
