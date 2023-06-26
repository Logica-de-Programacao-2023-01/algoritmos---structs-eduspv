package main

import "fmt"

// Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço".
// Escreva uma função que receba um slice de Viagens como parâmetro e retorne a viagem mais cara.
type viagem struct {
	origem  string
	destino string
	data    string
	preço   float64
}

func viagemMaisCara(viagens []viagem) viagem {
	var viagemMaisCara viagem
	precoMaximo := 0.0

	for _, v := range viagens {
		if v.preço > precoMaximo {
			precoMaximo = v.preço
			viagemMaisCara = v
		}
	}
	return viagemMaisCara
}
func main() {
	v1 := viagem{
		origem:  "brasília",
		destino: "são paulo",
		data:    "19 de abril de 2024",
		preço:   327.90,
	}
	v2 := viagem{
		origem:  "brasília",
		destino: "goiania",
		data:    "20 de abril de 2024",
		preço:   180.90,
	}
	v3 := viagem{
		origem:  "brasilia",
		destino: "new york",
		data:    "28 de fevereiro de 2030",
		preço:   1000.99,
	}
	v := []viagem{v1, v2, v3}
	ViagemMaisCara := viagemMaisCara(v)
	fmt.Printf("A viagem mais cara é de %s para %s, ocorrendo no dia %s com preço de R$ %.2f\n", ViagemMaisCara.origem, ViagemMaisCara.destino, ViagemMaisCara.data, ViagemMaisCara.preço)
}
