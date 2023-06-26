package main

import "fmt"

// Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas".
// O campo "notas" deve ser um slice de float64, representando as notas que o aluno tirou em uma determinada disciplina.
// Escreva funções que permitam adicionar ou remover notas do aluno, calcular a média das notas e imprimir o nome, idade e média do aluno.
type aluno struct {
	nome  string
	idade int
	notas []float64
}

func adicionarNota(aluno1 aluno) aluno {
	nota := 0.0
	for {
		adicionar := ""
		fmt.Println("digite a nota do seu aluno")
		fmt.Scan(&nota)
		aluno1.notas = append(aluno1.notas, nota)
		fmt.Println("Voce deseja adicionar outra nota?(n/s)")
		fmt.Scan(&adicionar)
		if adicionar == "n" {
			break
		}
	}
	return aluno1
}
func removerNota(aluno1 aluno) aluno {
	for {
		removerOutra := ""
		indexDaNotaRemovida := 0
		fmt.Println(aluno1.notas)
		fmt.Println("qual das notas voce deseja remover (começando de 0 terminando no 3)")
		fmt.Scan(&indexDaNotaRemovida)
		aluno1.notas = append(aluno1.notas[:indexDaNotaRemovida], aluno1.notas[indexDaNotaRemovida+1:]...)
		fmt.Println("deseja remover outra nota?(n/s)")
		fmt.Scan(&removerOutra)
		if removerOutra == "n" {
			break
		}
	}
	return aluno1
}
func MediaNotas(aluno1 aluno) {
	MediaAritimetica := 0.0
	Soma := 0.0
	for _, media := range aluno1.notas {
		Soma += media
	}
	MediaAritimetica = Soma / float64(len(aluno1.notas))
	fmt.Println(MediaAritimetica)
}
func printDados(aluno1 aluno) {
	fmt.Println(aluno1.nome)
	fmt.Println(aluno1.idade)
	fmt.Println(aluno1.notas)
}
func main() {
	aluno1 := aluno{
		nome:  "eduardo",
		idade: 18,
		notas: []float64{},
	}
	aluno1 = adicionarNota(aluno1)

	remover := ""
	fmt.Println("voce deseja remover alguma nota?(s/n)")
	fmt.Scan(&remover)
	if remover == "s" {
		aluno1 = removerNota(aluno1)
	}
	MediaNotas(aluno1)
	printDados(aluno1)
}
