package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda Feira"
	case 3:
		return "Terça Feira"
	case 4:
		return "QUarta feira"
	case 5:
		return "Quinta Feira"
	case 6:
		return "Sexta Feira"
	case 7:
		return "Sabado Feira"
	default:
		return "Numero invalido"
	}

}

func main() {
	fmt.Println("Switch")
	
	fmt.Println("---------------")
	dia := diaDaSemana(13)
	fmt.Println(dia)
}
