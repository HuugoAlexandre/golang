package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// Criação de arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	
	tamanho, err := f.WriteString("Olá, mundo!")
	if err != nil {
		panic(err)
	}
	fmt.Println(tamanho, "bytes escritos com sucesso!")
	f.Close()
	
	// Leitura de arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura de arquivo com buffer
	archive, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(archive)
	buffer := make([]byte, 6)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}