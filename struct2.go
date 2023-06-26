package main

import "fmt"

// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço".
// O campo "endereço" deve ser uma outra struct
// com os campos "rua", "número", "cidade" e "estado".
// Escreva uma função que receba uma Pessoa como parâmetro
// e imprima seu endereço completo.
type pessoa struct {
	nome     string
	idade    string
	endereco dados
}
type dados struct {
	rua    string
	numero string
	cidade string
	estado string
}

func EnderecoCompleto(p pessoa) string {
	return p.nome + p.idade + p.endereco.cidade + p.endereco.rua + p.endereco.estado + p.endereco.numero
}
func main() {
	p1 := pessoa{
		nome:  "eduardo ",
		idade: "18 ",
		endereco: dados{
			rua:    "octogonal ",
			numero: "315 ",
			cidade: "brasilia ",
			estado: "DF ",
		},
	}
	resultado := EnderecoCompleto(p1)
	fmt.Println(resultado)
}
