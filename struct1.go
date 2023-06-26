package main

import (
	"fmt"
	"math"
)

// Crie uma struct chamada Circulo com o campo "raio".
// Escreva uma função que receba um Circulo como parâmetro e
// calcule a área do círculo (área = pi * raio * raio).
// Dica: use a constante math.Pi para representar o número pi.
type circulo struct {
	raio float64
}

func AreaCirculo(c1 circulo) float64 {
	return math.Pi * c1.raio * c1.raio
}

func main() {
	c1 := circulo{raio: 0.5}
	resultado := AreaCirculo(c1)
	fmt.Println(resultado)

}
