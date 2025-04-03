package main

import "fmt"

func main(){
var ages = [4]int{17, 16, 20, 40}
nomes := [4]string{"Amanda", "Ana", "Julia", "Lorena"}
fmt.Println(ages)
fmt.Println(nomes)
nomes[3] = "Roberto"
fmt.Println(ages)
// Slice
var score =[]int{100, 200, 300, 400}
fmt.Println(score)
score[1] = 2
fmt.Println(score, len(score), cap(score))
rangeOne := score[1:3]
fmt.Println(rangeOne)
rangeTwo := score[2:]
fmt.Println(rangeTwo)
rangeThree := score[3]
fmt.Println(rangeThree)

var amigos = []string{"Ana", "Julia", "Lorena"}
fmt.Println(amigos)
amigos = append(amigos, "isa", "yas", "manu", "heitor")
fmt.Println(amigos, len(amigos), cap(amigos))

}                