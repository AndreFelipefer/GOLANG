package main

import "fmt"

func main() {
	var variavel1 string = "Testando variavel"
	variavel2 := "testando variavel oculta"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "testando outro modo de declarar variavel"
		variavel4 string = "acho que deu certo"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)
}