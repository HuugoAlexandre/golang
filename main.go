// ARQUIVO COM SINTAXE BÁSICA DE GO
package main

import (
	"errors"
	"fmt"
)

// STRUCT
type Endereco struct {
	logradouro string
	numero int
	cidade string
	estado string
}

type Carteirinha struct {
	id int
	time string
	anoDeEntrada int
}

type Cliente struct {
	nome string
	idade int
	ativo bool
	// Composição anônima em Struct
	Endereco
	// Composição nomeada
	carteira Carteirinha
}

func (c *Cliente) desativar() {
	c.ativo = false
	fmt.Printf("\nO clinte %s foi desativado.", c.nome)
}

func main() {

	// ARRAYS
	var meuArray [3]int
	meuArray[0] = 5
	meuArray[1] = 6
	meuArray[2] = 7
	fmt.Printf("Primeiro elemento: %d", meuArray[0])
	fmt.Printf("\nSegundo elemento: %d", meuArray[1])
	fmt.Printf("\nÚltimo elemento: %d\n", meuArray[len(meuArray)-1])
	// SLICES
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[4:]), cap(slice[4:]), slice[:4])

	// O slice opera com arrays estáticos nos bastidores,
	// criando um novo array estático com o dobro da capacidade
	// É melhor definir um array estático se souber mensurar a quantidade de elementos,
	// evitando a sobrecarga de Golang ter que redimensioná-lo frequentemente.
	slice = append(slice, 8)
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])

	// MAPS
	salarios := map[string]int{"hugo": 1200, "luan": 2400, "arthur": 3600}
	fmt.Println(salarios["luan"])
	fmt.Println(salarios)
	delete(salarios, "hugo")
	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Printf("Nome: %s, Salario: %d\n", nome, salario)
	}

	// blank identifier
	for _, salario := range salarios {
		fmt.Printf("Salario: %d\n", salario)
	}

	// outra forma de criar maps
	infos := make(map[string]int)
	infos["victor"] = 19
	infos["klinger"] = 30
	infos["joao"] = 17
	for nome, idade := range infos {
		fmt.Printf("\nNome: %s, Idade: %d", nome, idade)
	}

	fmt.Println(soma(2,3))
	fmt.Println(produto(5,21))
	fmt.Println(somaMaiorQueCem(50,49))
	fmt.Println(somaMaiorQueCem(50,51))
	// função que retorna 2 tipos, caso queira ser associada, precisa de 2 variáveis
	sum, err := somaMaiorQueCem(5, 25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)

	somaParametrosVariaveis := somaXParametros(2,5,6,7)
	fmt.Println(somaParametrosVariaveis)

	// Funções anônimas (closures), funções dentro de funções
	totalMultiplicado := func () int {
		return somaXParametros(2,5,6,7) * 2
	}()

	fmt.Println(totalMultiplicado)

	hugo := Cliente {
		nome: "Hugo",
		idade: 21,
		ativo: true,
	}
	fmt.Printf("Nome: %s, idade: %d, ativo: %t", hugo.nome, hugo.idade, hugo.ativo)
	hugo.ativo = true
	hugo.cidade = "roteiro" // mesma coisa que hugo.Endereco.cidade = "roteiro"
	hugo.carteira.time = "Corinthians" // composição nomeada precisa incluir o campo da struct
	fmt.Printf("\nNome: %s, idade: %d, ativo: %t, cidade: %s, time: %s", hugo.nome, hugo.idade, hugo.ativo, hugo.cidade, hugo.carteira.time)
	hugo.desativar()
	fmt.Printf("\nNome: %s, idade: %d, ativo: %t, cidade: %s, time: %s", hugo.nome, hugo.idade, hugo.ativo, hugo.cidade, hugo.carteira.time)
}

// FUNÇÕES
func soma(a int, b int) int {
	return a + b
}

// Função pode retonar dois valores
func produto(a, b int) (int, bool) {
	if a * b >= 100 {
		return a * b, true
	} else {
		return a * b, false
	}
}

// Isso é muito usado pra tratar erro
func somaMaiorQueCem(a, b int) (int, error) {
	if a + b > 100 {
		return 0, errors.New("A soma ultrapassou o valor 100.")
	} 
	// nil é vazio, nulo
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
