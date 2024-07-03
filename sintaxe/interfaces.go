package main

// INTERFACE
type Pessoa interface {
	ativar()
}

type Empresa struct {
	nome string
}

func (e Empresa) ativar() {}

func ativacao(pessoa Pessoa) {
	pessoa.ativar()
}
