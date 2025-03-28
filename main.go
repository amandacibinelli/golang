package main

import "fmt"

func main() {
var usuario string
var senha int
fmt.Println ("Digte seu usuario")
fmt.Scan(&usuario)
if usuario == "admin" {
fmt.Println("Digite sua senha")
fmt.Scan(&senha) 
if senha == 1234 {
fmt.Println("acesso permitido")
} else {
fmt.Println("senha incorreta")
}
} else { 
fmt.Println("usuario incorreto")
}

}