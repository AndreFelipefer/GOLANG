package main

import (
	"fmt"
)

func p(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	p("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario1 := map[string]map[string]string{
		"nome" :{
			"First Name": "Luiz",
			"Last Name" : "Otavio",
		},
		"Curso": {
			"Cursando": "ADS",
			"Faculdade": "Faceens",
		},
	}

	fmt.Println(usuario1)
	delete(usuario1, "nome")
	fmt.Println(usuario1)

	usuario1["signo"] = map[string]string{
		"Nome" : "Aquariano",
	}
	fmt.Println(usuario1)

	usuario2 := map[string]string{
		"Nome: ": "Gemeos",
	}
	fmt.Println(usuario2)
}
