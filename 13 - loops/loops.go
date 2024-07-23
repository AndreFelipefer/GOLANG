package main

import (
	"fmt"
	// "time"
)

func main() {

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando I - " , i)
	// 	time.Sleep(time.Second)

	// }

	// fmt.Println(i)

	// for j := 0 ; j < 10; j++{
	// 	fmt.Println("Incrementando j - ", j)
	// 	time.Sleep(time.Second)
	// }

	// array de string com noems
	nomes := [3]string{"Joao", "Davi", "Lucas"}

	// Looping de repetição e interação exibindo indice, nome por todo o array
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Looping de repetição exibindo cada letra de acordo com  o indice na frase
	for indice, letra := range "PALAVRA" {
		// exibição do indice mais a letra formatada da tabela ask
		fmt.Println(indice, string(letra))
	}

	// Criando map com informações do usuario
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	// Looping de repetição exibindo no formato map a chave e o valor
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	
}
