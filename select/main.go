package main

import (
	"fmt"
	"time"
)

// Select funciona como um Switch para canais
// Select permite guardar operações de vários canais

func main() {
	// ================  Criação de canais
	// Aqui, dois canais são criados, channel1 e channel2, ambos capazes de
	// transmitir valores do tipo string. Os canais em Go são uma
	// forma de comunicação entre goroutines.
	channel1 := make(chan string)
	channel2 := make(chan string)

	// ================  Lançamento de goroutines
	// Esta goroutine (go func()) espera 1 segundo (time.Sleep(time.Second * 1))
	// antes de enviar a string "one" para o canal channel1 usando o operador de envio <-.
	go func() {
		time.Sleep(time.Second * 1)
		channel1 <- "one"
	}()

	// Outra goroutine é iniciada, mas desta vez espera 2 segundos
	// antes de enviar a string "two" para o canal channel2.
	go func() {
		time.Sleep(time.Second * 2)
		channel2 <- "two"
	}()

	// Essas goroutines são executadas concorrrentemente com o restante do programa.

	// O loop for é configurado para iterar duas vezes, já que o programa
	// espera receber mensagens de dois canais.
	for i := 0; i < 2; i++ {

		// ======= Funcionamento do select:
		// O select funciona como um switch para canais. Ele verifica se
		// alguma operação de leitura ou escrita nos canais está pronta.
		// Ele bloqueia a execução até que uma das operações esteja disponível.
		// Assim que uma operação é completada, o respectivo bloco de código é executado.
		select {
		case msg1 := <-channel1:
			fmt.Println("Recebeu", msg1)
		case msg2 := <-channel2:
			fmt.Println("Recebeu", msg2)
		}
	}

}

// O uso do select permite que o programa reaja a eventos de canais
// assim que eles acontecem, sem a necessidade de verificar cada canal
// individualmente ou usar goroutines adicionais para monitoramento.
// Isso é útil em sistemas concorrentes que precisam lidar com múltiplos
// fluxos de dados.
