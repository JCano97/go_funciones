package main

import "fmt"

func fibonacci(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	fmt.Print("Ingrese la cantidad de numeros a mostrar de la serie fibonacci: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(uint(i)), ",")
	}
}
