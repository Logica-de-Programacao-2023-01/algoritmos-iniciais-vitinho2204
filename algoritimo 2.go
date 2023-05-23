package main

import "fmt"

//Algoritmo de conversão de temperatura: um algoritmo
//que converte uma temperatura em Celsius para Fahrenheit ou vice-versa.
//multiplicar a temperatura em graus Celsius por 1,8 e somar 32

func main() {
	var C int
	fmt.Print(" degite a temperatura em graus Celsius ")
	fmt.Scan(&C)
	fmt.Print(" a temperatura em graus Fahrenheit é: ", C*18/10+32)

}
