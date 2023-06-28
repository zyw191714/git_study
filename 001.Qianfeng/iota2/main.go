package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
		D = 100
		E = 101
		F = iota
		G
	)
	const H = iota
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
}
