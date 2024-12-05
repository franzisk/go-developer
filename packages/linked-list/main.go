package main

import (
	"container/list"
	"fmt"
)

func main() {

	// Cria uma nova lista e adiciona alguns números
	list := list.New()

	// Adiciona o valor 4 ao final da lista.
	// Retorna um ponteiro para o elemento inserido (e4), que pode ser
	// usado para realizar operações subsequentes como inserir antes ou depois dele.
	e4 := list.PushBack(4)

	// Adiciona o valor 1 ao início da lista.
	// Retorna um ponteiro para o elemento inserido (e1).
	e1 := list.PushFront(1)

	// Insere o valor 3 antes do elemento e4 (que contém 4).
	list.InsertBefore(3, e4)

	//Insere o valor 2 depois do elemento e1 (que contém 1).
	list.InsertAfter(2, e1)

	list.InsertAfter(12, e1)
	list.InsertAfter(13, e1)
	list.InsertAfter(14, e1)

	// Interage na lista e imprime o conteúdo
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
