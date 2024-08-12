package main

import (
	"fmt"
	"teste/sintaxe/formula1"
)

func main() {
	// ARRAYS
	meuArray := [3]int{5, 6, 7}
	fmt.Printf("Primeiro elemento: %d\n", meuArray[0])
	fmt.Printf("Segundo elemento: %d\n", meuArray[1])
	fmt.Printf("Último elemento: %d\n", meuArray[len(meuArray)-1])

	// SLICES
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[:2]), cap(slice[:2]), slice[:2])
	fmt.Printf("len=%d  cap=%d  %v\n", len(slice[4:]), cap(slice[4:]), slice[:4])
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

	for _, salario := range salarios {
		fmt.Printf("Salario: %d\n", salario)
	}

	infos := make(map[string]int)
	infos["victor"] = 19
	infos["klinger"] = 30
	infos["joao"] = 17
	for nome, idade := range infos {
		fmt.Printf("\nNome: %s, Idade: %d", nome, idade)
	}

	fmt.Println(soma(2, 3))
	fmt.Println(produto(5, 21))
	fmt.Println(somaMaiorQueCem(50, 49))
	fmt.Println(somaMaiorQueCem(50, 51))

	sum, err := somaMaiorQueCem(5, 25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)

	somaParametrosVariaveis := somaXParametros(2, 5, 6, 7)
	fmt.Println(somaParametrosVariaveis)

	totalMultiplicado := func() int {
		return somaXParametros(2, 5, 6, 7) * 2
	}()
	fmt.Println(totalMultiplicado)

	hugo := Cliente{
		nome: "Hugo",
		idade: 21,
		ativo: true,
		Endereco: Endereco{
			cidade: "roteiro",
		},
		carteira: Carteirinha{
			time: "Corinthians",
		},
	}
	fmt.Printf("\nNome: %s, idade: %d, ativo: %t, cidade: %s, time: %s", hugo.nome, hugo.idade, hugo.ativo, hugo.cidade, hugo.carteira.time)
	hugo.desativar()
	fmt.Printf("\nNome: %s, idade: %d, ativo: %t, cidade: %s, time: %s", hugo.nome, hugo.idade, hugo.ativo, hugo.cidade, hugo.carteira.time)
	ativacao(&hugo)
	fmt.Printf("\nNome: %s, idade: %d, ativo: %t, cidade: %s, time: %s", hugo.nome, hugo.idade, hugo.ativo, hugo.cidade, hugo.carteira.time)
	minhaEmpresa := Empresa{}
	ativacao(minhaEmpresa)
	fmt.Println()
	// POINTERS
	valorA := 5
	valorB := 6
	resultadoMulti := multi(&valorA,&valorB)
	fmt.Println(resultadoMulti)
	fmt.Println(valorA, valorB)

	// Exemplos de uso de ponteiro com struct
	hugoConta := Conta{saldo: 500}
	hugoConta.deposita(200)
	fmt.Println(hugoConta.saldo)

	hugoNovaConta := novaConta()
	fmt.Println(hugoNovaConta.saldo)
	hugoNovaConta.deposita(300)
	fmt.Println(hugoNovaConta.saldo)

	// type assertion
	var minhaVar interface{} = 12 // interface vazia
	fmt.Println(minhaVar)
	res, ok := minhaVar.(int)
	if ok {
		fmt.Printf("minhaVar é inteiro: %v", res)
	} else {
		fmt.Printf("minhaVar NÃO é inteiro: %v", res)
	}
	
	// Generics
	fmt.Println()
	mapa := map[string]int {"hugo": 21, "luan": 20, "arthur": 19}
	mapa2 := map[string]float64 {"hugo": 21.5, "luan": 20.5, "arthur": 19.5}
	// coloca ~ para indicar que meuNumero é um tipo genérico que é do tipo int explicitamente
	mapa3 := map[string]meuNumero {"hugo": 21, "luan": 20, "arthur": 19}
	fmt.Println(somaMapa(mapa))
	fmt.Println(somaMapa2(mapa2))
	fmt.Println(somaMapa2(mapa3)) 

	// constraint comparable
	fmt.Println(comparar(5, 5.0))

	// Teste do uso de go.mod
	multiPackage := formula1.Multi(5, 7)
	fmt.Println(multiPackage)
	
	// cor está em letra minúscula em sua definição, não pode ser acessada
	lewis := formula1.Montadora{Nome: "Mercedes"}
	fmt.Println(lewis.Nome)

}

