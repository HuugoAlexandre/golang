package main

import "fmt"

// STRUCT
type Endereco struct {
	logradouro string
	numero     int
	cidade     string
	estado     string
}

type Carteirinha struct {
	id          int
	time        string
	anoDeEntrada int
}

type Cliente struct {
	nome     string
	idade    int
	ativo    bool
	Endereco // Composição anônima em Struct
	carteira Carteirinha // Composição nomeada
}

func (c *Cliente) desativar() {
	c.ativo = false
	fmt.Printf("\nO cliente %s foi desativado.", c.nome)
}

func (c *Cliente) ativar() {
	c.ativo = true
	fmt.Printf("\nO cliente %s foi ativado.", c.nome)
}
