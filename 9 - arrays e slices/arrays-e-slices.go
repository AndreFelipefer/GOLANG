package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array  Slices")

	var array1 [5]int
	array1[0] = 3
	fmt.Println(array1)

	array2 := [5]string{
		"posição 1",
		"posição 2",
		"posição 3 ",
		"posição 4",
		"posição 5",
	}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	// Criação de slice
	slice := []int{10, 11, 13, 14, 16}

	// Alterando o array de acordo com o indice
	slice[0] = 1

	//Adicionando slice utilizando append
	slice = append(slice, 10)

	// Printando Slice
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(len(slice3)) // verificar tamanho (LENGHT)
	fmt.Println(cap(slice3)) // Capacidade 

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 7)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // verificar tamanho (LENGHT)
	fmt.Println(cap(slice3)) // Capacidade 


	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
