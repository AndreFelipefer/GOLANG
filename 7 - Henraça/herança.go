package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso       string
	instituição string
	ra          uint32
}

func main() {
	fmt.Println("Henraca")

	aluno := pessoa{"Jorge", 28, "Camargo"}
	// fmt.Println(aluno)

	aluno1 := estudante{aluno, "ADS", "Facens", 404040}
	fmt.Println(aluno1)

}
