package main

import (
	"fmt"
	"go-developer/utils"
	"os"
)

type Pessoa struct {
	nome  string
	idade int
}

func calcularMedia(listaNotas []float64) float64 {
	total := 0.0
	for _, nota := range listaNotas {
		total += nota
	}
	media := total / float64(len(listaNotas))
	return media
}

// Exemplo de closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Exemplo de recursão
func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

// O defer em Go é utilizado para atrasar a execução de uma função até que a
// função que o contém termine. Isso é útil para garantir a execução de tarefas como
// fechar arquivos, liberar recursos ou registrar informações de depuração.
func deferExample() {
	fileName := "example.txt"

	// Abrir o arquivo
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}

	// Garantir o fechamento do arquivo após terminar a execução
	// O defer file.Close() é executado no início, mas a execução
	// de file.Close() é atrasada até que a função deferExample() termine.
	defer file.Close()

	// Simular a leitura do arquivo
	fmt.Printf("Arquivo %s aberto com sucesso!\n", fileName)
	// Aqui poderiam ser feitas operações com o arquivo, como leitura/escrita
}

func panicExample() {
	// Antes de fechar/finalizar a função o defer recolhe o erro ocorrido e mostra na tela
	defer func() {
		// a chamada a recover() captura o valor passado para panic
		// "recover" funciona somente dentro de um "refer"
		x := recover()
		utils.Print("ERRO:", x)
	}()

	// O programa interrompe a execução normal da função
	panic("O cadastro não foi encontrado no sistema")
}

// Exemplo usando ponteiros
// Função que altera o valor de uma variável usando um ponteiro
func dobrar(valor *int) {
	*valor = *valor * 2 // Acessa o valor apontado pelo ponteiro e dobra
}

// Função que altera a idade de uma pessoa
func fazerAniversario(p *Pessoa) {
	p.idade++ // Modifica diretamente o campo 'idade'
}

func main() {
	listaNotas := []float64{7.7, 6.8, 9.9, 8, 75}
	media := calcularMedia(listaNotas)
	utils.Print("Média: ", media)

	nextInt := intSeq()
	utils.Print(nextInt())
	utils.Print(nextInt())
	utils.Print(nextInt())

	utils.Print(fatorial(9))

	deferExample()

	// Correção: Envolver a chamada a `fatorial(9)` em uma função anônima
	utils.ExecWithRecover(func() {
		result := fatorial(9)
		fmt.Printf("Fatorial de 9 é %d\n", result)
	})

	panicExample()

	// EXEMPLO DE USO DE PONTEIROS
	x := 10
	// Passa o endereço da variável para a função, não valor propriamente
	dobrar(&x)

	pessoa := Pessoa{nome: "João", idade: 30}
	fmt.Println("Antes do aniversário:", pessoa)

	// Passa o ponteiro da estrutura
	fazerAniversario(&pessoa)
	fmt.Println("Depois do aniversário:", pessoa)
}
