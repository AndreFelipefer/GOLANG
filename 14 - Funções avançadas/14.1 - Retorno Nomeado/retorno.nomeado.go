package main

import "fmt"

// Retorno nomeado facilita o entendimento para um código mais limpo

// Função com retorno nomeado  -- declaração de variavel e o tipo
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	// variavel sem a inferencia pois o tipo foi informado na propria funçãoa acima
	soma = n1 + n2
	subtracao = n1 - n2
	//retorna os calculos
	return
}

func main() {
	// declarando as variaveis e chamando a função e atribuindo valores a serem calculados
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
