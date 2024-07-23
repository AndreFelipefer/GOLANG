package main

import "fmt"

func main() {
	var variavel1 int8 = 10
	var variavel2 int8 = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	// DESREFERENCIAÇÃO 
	fmt.Println(variavel3, *ponteiro)



}
