package main

func multi(a, b *int) int {
	// Muda o valor guardado no endereço apontado pelo ponteiro
	*a = 40
	*b = 70
	return *a * *b
}

// Exemplos de uso de ponteiros com struct
type Conta struct {
	saldo int 
}

func (c *Conta) deposita(valor int) int {
	c.saldo += valor
	return c.saldo
}

// Criação de uma nova conta
func novaConta() *Conta {
	return &Conta{saldo: 0}
}