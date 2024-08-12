package formula1

// Funções, métodos, interfaces etc que começam com letra maiúscula 
// podem ser exportadas para outros pacotes
func Multi[T ~int | ~float64](a, b T) T {
	return a * b
}

var H string = "Hugo"

type Pessoa struct {
	Nome string
}

// Montadora pode ser encontrada em outros pacotes
// mas seu atributo cor não pode ser acessado
type Montadora struct {
	Nome string
	cor string
}