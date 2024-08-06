package main

// Generics
func somaMapa[T ~int | ~float64](m map[string]T) T {
	var total T = 0
	for _, v := range m {
		total += v
	}
	return total
}

// Também pode ser construído com constraint
type Number interface {
	~int | ~float64
}

func somaMapa2[T Number](m map[string]T) T {
	var total T = 0
	for _, v := range m {
		total += v
	}
	return total
}

// constraint comparable compara igualdade entre dois valores
func comparar[T comparable](a, b T) bool {
	return a == b
}
type meuNumero int

