package main

import "fmt"

//Algoritmo de cálculo de média: um algoritmo
//que calcula a média de uma lista de números.

func main() {
	var n1 int
	var n2 int
	var n3 int
	var n4 int
	fmt.Print(" numero 1: ")
	fmt.Scan(&n1)
	fmt.Print(" numero 2: ")
	fmt.Scan(&n2)
	fmt.Print(" numero 3 :")
	fmt.Scan(&n3)
	fmt.Print(" numero 4: ")
	fmt.Scan(&n4)
	fmt.Print(" a media é ", n1/4+n2/4+n3/4+n4/4)

}
