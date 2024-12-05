package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding"
	"fmt"
	"hash/crc32"
	"log"
)

func test1() {
	const (
		input1 = "A marmorta cava fundo para baixo, "
		input2 = "sem nenhuma ideia do que vai encontrar."
	)

	// Cria uma nova instância do hash SHA-256.
	primeiro := sha256.New()

	// Escreve os bytes da string input1 no hash primeiro.
	primeiro.Write([]byte(input1))

	//  Faz uma type assertion para verificar se primeiro implementa a interface BinaryMarshaler.
	// A interface é usada para converter o estado do hash em uma sequência de bytes.
	marshaler, ok := primeiro.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("primeiro não implementa encoding.BinaryMarshaler")
	}

	// Serializa o estado interno do hash primeiro em um slice de bytes (state).
	// Isso permite salvar o progresso do cálculo.
	state, err := marshaler.MarshalBinary()

	if err != nil {
		log.Fatal("Unable to marshal hash", err)
	}

	// Cria uma nova instância do hash SHA-256.
	segundo := sha256.New()

	// Faz uma type assertion para verificar se segundo implementa a interface
	// BinaryUnmarshaler. Essa interface permite restaurar o estado interno a partir de um slice de bytes.
	unmarshaler, ok := segundo.(encoding.BinaryUnmarshaler)

	if !ok {
		log.Fatal("o segundo não implementa encoding.BinaryUnmarshaler")
	}

	// Restaura o estado do hash segundo com o estado previamente serializado (state).
	if err := unmarshaler.UnmarshalBinary(state); err != nil {
		log.Fatal("Unable to unmarshal hash: ", err)
	}

	// Continua escrevendo os bytes de input2 no hash primeiro.
	primeiro.Write([]byte(input2))

	// Escreve os mesmos bytes no hash segundo
	segundo.Write([]byte(input2))

	// Como segundo foi restaurado para o mesmo estado que primeiro após o
	// cálculo de input1, ambos hashes devem produzir o mesmo resultado final.

	// Calcula o valor final do hash para primeiro.
	// Imprime o hash resultante em formato hexadecimal.
	fmt.Printf("%x\n", primeiro.Sum(nil))

	// Compara os valores finais dos hashes primeiro e segundo. Como ambos processaram
	// exatamente os mesmos dados, incluindo a restauração do estado, o resultado será true.
	fmt.Println(bytes.Equal(primeiro.Sum(nil), segundo.Sum(nil)))

}

func test2() {
	// criar o hash
	h := crc32.NewIEEE()

	// escrever um dado no hash
	h.Write([]byte("texto para ser colocado no hash"))

	v := h.Sum32()

	fmt.Println(v)
}

func test3() {

	h := sha1.New()

	h.Write([]byte("texto para ser criptografado no SHA1"))

	bs := h.Sum([]byte{})

	fmt.Println(bs)

	fmt.Printf("%x\n", bs)
}

func main() {
	test1()
	test2()
	test3()
}
