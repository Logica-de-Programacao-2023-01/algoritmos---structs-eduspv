package main

import "fmt"

type animal struct {
	nome    string
	espécie string
	idade   int
	som     string
}

func ModificarSom(a animal, novoSom string) animal {
	a.som = novoSom
	return a
}

func PrintAnimal(a animal) {
	fmt.Printf("Nome: %s\n", a.nome)
	fmt.Printf("Espécie: %s\n", a.espécie)
	fmt.Printf("Idade: %d\n", a.idade)
	fmt.Printf("Som: %s\n", a.som)
}

func main() {
	a := animal{
		nome:    "leão",
		espécie: "leão",
		idade:   5,
		som:     "roaaar",
	}

	PrintAnimal(a)

	a = ModificarSom(a, "rugido")
	fmt.Println("Após modificar o som:")
	PrintAnimal(a)
}
