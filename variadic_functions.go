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
	a := []int{1, 4, 2, -1}
	fmt.Println(sum(4, 5, 2, 0, 8, 3, 7, 6))
	fmt.Println(sum(a...))
}
