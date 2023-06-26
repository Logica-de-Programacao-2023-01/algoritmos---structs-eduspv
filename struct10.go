package main

import "fmt"

// Crie uma struct chamada Filme com os campos "título", "diretor", "ano" e "avaliações".
// O campo "avaliações" deve ser um slice de inteiros representando as notas que o filme recebeu dos espectadores.
// Escreva funções que permitam adicionar ou remover avaliações do filme, calcular a média das avaliações e imprimir
// as informações do filme e sua média de avaliações.
type filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func adicionaravaliaçoes(filme1 filme) filme {
	avaliação := 0
	for {
		OutraAvaliação := ""
		fmt.Println("adicione uma avaliação")
		fmt.Scan(&avaliação)
		filme1.avaliacoes = append(filme1.avaliacoes, avaliação)
		fmt.Println("deseja adicionar uma nova avaliação?(s/n)")
		fmt.Scan(&OutraAvaliação)
		if OutraAvaliação == "n" {
			break
		}
	}
	return filme1
}
func removerAvaliações(filme1 filme) filme {
	for {
		removerOutra := ""
		indexDaNotaRemovida := 0
		fmt.Println(filme1.avaliacoes)
		fmt.Println("qual das notas voce deseja remover (começando de 0 terminando no 3)")
		fmt.Scan(&indexDaNotaRemovida)
		filme1.avaliacoes = append(filme1.avaliacoes[:indexDaNotaRemovida], filme1.avaliacoes[indexDaNotaRemovida+1:]...)
		fmt.Println("deseja remover outra nota?(n/s)")
		fmt.Scan(&removerOutra)
		if removerOutra == "n" {
			break
		}
	}
	return filme1
}
func CalcularMedia(filme1 filme) {
	MediaAritimetica := 0
	Soma := 0
	for _, media := range filme1.avaliacoes {
		Soma += media
	}
	MediaAritimetica = Soma / len(filme1.avaliacoes)
	fmt.Println("a media arrendondada é de:", MediaAritimetica)
}
func printfilme(filme1 filme) {
	fmt.Println(filme1.titulo)
	fmt.Println(filme1.ano)
	fmt.Println(filme1.diretor)
	fmt.Println(filme1.avaliacoes)
}
func main() {
	filme1 := filme{
		titulo:     "prenda-me se for capaz",
		diretor:    "Steven Spielberg",
		ano:        2003,
		avaliacoes: []int{},
	}
	filme1 = adicionaravaliaçoes(filme1)

	remover := ""
	fmt.Println("deseja remover alguma avaliação(s/n)")
	fmt.Scan(&remover)
	if remover == "s" {
		filme1 = removerAvaliações(filme1)
	}
	CalcularMedia(filme1)
	printfilme(filme1)
}
