package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int32 = 1000000
	fmt.Println(numero)

	var numero2 uint32 = 1234
	fmt.Println(numero2)

	// alias
	// int 32 = rune
	var numero3 rune = 1235
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// reais
	var numero5 float32 = 13.99
	var numero6 float64 = 12.99
	fmt.Println(numero5, numero6)

	// fim numeros reais

	// string
	var str string = "cadeia de caracteres"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)
	// fim string

	// booleano
	var booleano bool = true
	fmt.Println(booleano)

	//error
	var error error = errors.New("Erro de sintaxe")
	fmt.Println(error)

}
