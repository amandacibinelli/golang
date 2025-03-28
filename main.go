package main

import "fmt"

func main() {
	
	a , b := 10 , 3

	fmt.Println("A soma eh" , a + b)
	fmt.Println("A subtraçao: eh" , a - b)
	fmt.Println("A divisao eh" , a / b)
	fmt.Println("A multiplicaçao eh" , a * b)
	fmt.Println("O resto da divisao eh" , a % b)
a+=1
fmt.Println("Incrementar a " , a)

if a > 0 && b > 0 {
	fmt.Println("Numeros Positivos")
}
}