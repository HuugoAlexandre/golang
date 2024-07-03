package main

import "errors"

// FUNÇÕES
func soma(a int, b int) int {
	return a + b
}

// Função pode retornar dois valores
func produto(a, b int) (int, bool) {
	if a*b >= 100 {
		return a * b, true
	}
	return a * b, false
}

// Isso é muito usado pra tratar erro
func somaMaiorQueCem(a, b int) (int, error) {
	if a+b > 100 {
		return 0, errors.New("A soma ultrapassou o valor 100.")
	}
	return a + b, nil
}

// Funções variáticas (quantidade de parâmetros variáveis)
func somaXParametros(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
