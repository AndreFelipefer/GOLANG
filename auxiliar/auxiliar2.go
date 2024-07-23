package auxiliar

import "fmt"

// Escrevendo pela função 2
func Escrever2() {
	fmt.Println("Escrevendo pela função escrever2")
}

func p(args ...interface{}) {
	fmt.Println(args...)
}
