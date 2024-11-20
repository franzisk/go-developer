package utils

import "fmt"

// Print é uma função utilitária que imprime argumentos no console.
// Para ficar como exportável o nome deve iniciar com letra maiúscula
func Print(a ...any) {
	// desempacota os elementos do slice a e os passa individualmente para fmt.Println.
	fmt.Println(a...)
}

// ExecWithRecover executa a função fornecida e trata panics.
func ExecWithRecover(f func()) {
	// ativa o defer para recuperar o erro
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recuperado do pânico: %v\n", r)
		}
	}()

	// executa a função monitorada
	f()
}
