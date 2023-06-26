package main

import "fmt"

//Crie uma struct chamada Triângulo com os campos "base" e "altura".
//Escreva uma função que receba um Triângulo como parâmetro e calcule
//a área do triângulo (área = base * altura / 2).

type triangulo struct {
	base   int
	altura int
}

func AreaDoTriangulo(triangulo1 triangulo) int {
	area := (triangulo1.base * triangulo1.altura) / 2
	return area
}

func main() {
	triangulo1 := triangulo{base: 8, altura: 2}
	resultado := AreaDoTriangulo(triangulo1)
	fmt.Println(resultado)
}
