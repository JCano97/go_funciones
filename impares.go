package main

import "fmt"

func generadorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func main() {
	var rango int
	print("Ingrese el numero de impares a generar: ")
	fmt.Scan(&rango)

	siguienteImpar := generadorImpares()
	for i := 0; i < rango; i++ {
		fmt.Print(siguienteImpar(), ",")
	}
}
