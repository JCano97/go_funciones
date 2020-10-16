package main

import "fmt"

func intercambia(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var a int
	var b int
	fmt.Print("Ingresa el valor a: ")
	fmt.Scan(&a)
	fmt.Print("Ingresa el valor b: ")
	fmt.Scan(&b)

	intercambia(&a, &b)

	fmt.Println("Intercambiados")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
