package main

import "fmt"

func main() {

	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 2
	divisao := 10 / 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)


	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITMETICOS

	// ATRIBUIÇÃO

	var variavel1 string = "texto"
	variavel2 := "textinho basico"

	fmt.Println(variavel1, variavel2)

	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS 
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)


	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println("-------------------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM OPERADORES LÓGICOS

	// OPERADORES UNARIOS
	numero := 10
	numero ++
	numero += 15
	fmt.Println((numero))
	// FIM OPERADORES UNARIOS 

	
	var texto string
	if numero > 5 {
		texto = "maior que 5"
	}else {
		texto = "menor que 5"
	}

	fmt.Println(texto)
	
}



	