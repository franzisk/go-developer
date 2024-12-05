package main

import (
	"crypto/sha256"
	"encoding"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePath := "C:\\Users\\francisco.souza\\Projetos\\Java\\api-sistema-secsp\\target\\api-sistema-secsp.jar"
	chunkSize := 1024 * 1024 // Processar 1 MB por vez

	// Criar ou abrir o arquivo para simular o backup
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo:", err)
	}
	defer file.Close()

	// Criar o hash inicial
	hasher := sha256.New()

	// Ler uma parte do arquivo e processar
	buffer := make([]byte, chunkSize)
	bytesRead, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatal("Erro ao ler o arquivo:", err)
	}

	// Processar o primeiro chunk
	hasher.Write(buffer[:bytesRead])

	// Serializar o estado do hash após o primeiro chunk
	marshaler, ok := hasher.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("Hasher não suporta serialização")
	}

	hashState, err := marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("Erro ao serializar o estado do hash:", err)
	}

	// Simular pausa no processamento
	fmt.Println("Pausa no processamento...")
	fmt.Printf("Estado do hash salvo: %x\n", hasher.Sum(nil))

	// Simular retomada do processamento com um novo hash
	hasher2 := sha256.New()

	// Restaurar o estado do hash
	unmarshaler, ok := hasher2.(encoding.BinaryUnmarshaler)
	if !ok {
		log.Fatal("Hasher2 não suporta desserialização")
	}
	if err := unmarshaler.UnmarshalBinary(hashState); err != nil {
		log.Fatal("Erro ao restaurar o estado do hash:", err)
	}

	// Continuar lendo o arquivo e processando
	for {
		bytesRead, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal("Erro ao continuar a leitura:", err)
		}
		if bytesRead == 0 {
			break
		}
		hasher2.Write(buffer[:bytesRead])
	}

	// Exibir o hash final
	fmt.Printf("Hash final: %x\n", hasher2.Sum(nil))
}
