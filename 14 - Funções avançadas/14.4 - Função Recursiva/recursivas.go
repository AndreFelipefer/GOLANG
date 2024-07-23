package main

import "fmt"

func finobacci(posicao uint) uint{
	if posicao <= 1{
		return posicao
	}

	return finobacci(posicao - 2) + finobacci(posicao - 1)

}

func main() {
	fmt.Println("Funções Recursivas")
	//Formula de fibonacci
	 // 1 2  3  5 8 13 21 34 55

	posicao := uint(12)
	for i := uint(0); i < posicao; i++{

		fmt.Println(finobacci(i)) 
	}

}