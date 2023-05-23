package main

//Algoritmo de soma de dois números: um algoritmo simples que
//pede dois números como entrada e retorna a soma desses dois números.

import "fmt"

func main() {
	var n1 int
	fmt.Print("escreva o primeiro numero")

	fmt.Scan(&n1)
	var n2 int
	fmt.Print("escreva o segundo numero")

	fmt.Scan(&n2)
	fmt.Println(" a soma dos dois numeros é:", n1+n2)

}
