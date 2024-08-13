package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // tag para mudar o nome do campo
	Saldo int `json:"s"`
}

func main() {
	conta := Conta{123, 1000}
	resultado, err := json.Marshal(conta) 
	if err != nil {
		panic(err)
	}
	// json retorna em bytes, por isso é necessário converter para string
	fmt.Println(string(resultado))	
	
	// Com Encoder não precisa guardar em variável
	json.NewEncoder(os.Stdout).Encode(conta) 

	exJson := []byte(`{"n":456, "s":2000}`)
	var outraConta Conta
	err = json.Unmarshal(exJson, &outraConta) // faz o inverso do Marshal, converte bytes em struct
	if err != nil {
		panic(err)
	}
	fmt.Println(outraConta.Saldo)
}