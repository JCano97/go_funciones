package main

import "fmt"

func sum(args ...int) int {
	masGrande := 0
	for _, v := range args {
		if v > masGrande {
			masGrande = v
		}
	}
	return masGrande
}

func main() {
	var n int
	var entrada int

	fmt.Print("Ingrese la cantidad de enteros a capturar: ")
	fmt.Scan(&n)
	//define tamaño del slice
	a := make([]int, 0, n)
	//recibe los enteros
	for i := 0; i < n; i++ {
		fmt.Print("Ingrese entero: ")
		fmt.Scan(&entrada)
		a = append(a, entrada)
	}
	fmt.Println("Numero mas grande: ", sum(a...))
}
