package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função1")

}

func funcao2(){
	fmt.Println("Executando a função2")
	

}

func alunoEstaAprovado(n1, n2 float32) bool{
	fmt.Println("Entrando na função")
	// defer será exibido ao final do retorno da função 
	defer fmt.Println("Média calculada o resultado será retornado.")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}else{
		return false
	}
} 

func main() {
    // defer funcao1()  // Adia a execução de funcao1 até o retorno de main
    // funcao2()        // Executa funcao2 imediatamente
	fmt.Println(alunoEstaAprovado(7, 8))
}

