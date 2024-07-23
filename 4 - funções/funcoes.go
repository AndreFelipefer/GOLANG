package main

import (
	"fmt"
	
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(num1, num2 int8) (int8, int8){
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)


	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado  := f("Testando a função 1")
	fmt.Println(resultado)


	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)


}
