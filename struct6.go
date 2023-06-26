package main

import "fmt"

// Crie uma struct chamada Funcionário com os campos "nome",
// "salário" e "idade". Escreva funções que permitam aumentar ou diminuir o salário do funcionário
// em uma determinada porcentagem e uma função que calcule o tempo de serviço do funcionário
// (considerando que ele começou a trabalhar aos 18 anos)
type funcionarios struct {
	nome    string
	salario float32
	idade   int
}

func MudancaSalario(funcionarios funcionarios) {
	fmt.Printf("o salario com aumento ficaria %.2f\n", funcionarios.salario*1.10)
	fmt.Printf("o salario com a diminuicao ficaria %.2f\n", funcionarios.salario/1.10)
}
func TempoDeServico(funcionarios funcionarios) {
	fmt.Printf("O %s trabalha na empresa a %d anos", funcionarios.nome, funcionarios.idade-18)
}
func main() {
	funcionario := funcionarios{
		nome:    "Sergio",
		salario: 15000,
		idade:   30,
	}
	MudancaSalario(funcionario)
	TempoDeServico(funcionario)
}
