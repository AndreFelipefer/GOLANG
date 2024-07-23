package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("!")
	enderecoCompleto := endereco{"Rua do chaves", 0}
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{"davi", 21, enderecoCompleto}
	fmt.Println(u2)


	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)
}
