package main

import "fmt"

func main() {
var numero1 float32
var numero2 float32
fmt.Println ("Digte dois numero")
fmt.Scan(&numero1 , &numero2)
fmt.Println("O numero digitado eh:", numero1 , numero2)
	fmt.Println("A soma eh" , numero1 + numero2)
	fmt.Println("A subtraçao: eh" , numero1 - numero2)
	fmt.Println("A divisao eh" , numero1 / numero2)
	fmt.Println("A multiplicaçao eh" , numero1 * numero2)
}