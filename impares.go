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
	siguienteImpar := generadorImpares()
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
}
