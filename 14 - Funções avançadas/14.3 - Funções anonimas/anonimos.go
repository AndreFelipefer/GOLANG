package main

import "fmt"

func main() {
	// Função anonima sendo atribuida a variavel retorno com todos os paramentos inseridos dentro da própria função
	retorno := func(texto string) string {
		//Sprintf utilizado para concatenar e formatar tipos de dados 
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Paramêtro")

	fmt.Println(retorno)
}
