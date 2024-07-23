package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil{
		fmt.Println("Execução recuperada com sucesso")
	}

}

func calcularMedia(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// A função entra em  panico e finaliza a ação. o Panic sempre  recebe  argumento dentro dos parenteses
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(calcularMedia(8, 6))
	fmt.Println("Pós execução.")

}
