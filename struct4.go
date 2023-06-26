package main

import (
	"fmt"
	"time"
)

//Crie uma struct chamada Playlist com os campos "nome" e "músicas".
//O campo "músicas" deve ser um slice de structs, cada uma representando uma música
//com os campos "título", "artista" e "duração". Escreva uma função que receba uma Playlist
//como parâmetro e imprima o nome da playlist, o tempo total da playlist e as informações de cada música.

type playlist struct {
	nome    string
	musicas []struct {
		titulo  string
		artista string
		duracao float64
	}
}

func DadosPlaylist(playlist1 playlist) string {
	fmt.Printf("Nome da Playlist: %s\n", playlist1.nome)
	totalDuracao := 0.0
	for _, musica := range playlist1.musicas {
		fmt.Printf("Título: %s\n", musica.titulo)
		fmt.Printf("Artista: %s\n", musica.artista)
		fmt.Printf("Duração: %.2f\n", musica.duracao)
		totalDuracao += musica.duracao
		fmt.Println()
	}
	fmt.Printf("Tempo total da Playlist: %s\n", time.Duration(totalDuracao)*time.Minute)
	return ""
}
func main() {
	playlist1 := playlist{
		nome: "Deus é mais",
		musicas: []struct {
			titulo  string
			artista string
			duracao float64
		}{
			{
				titulo:  "eu sou da forca",
				artista: "forca jovem",
				duracao: 3.5,
			},
			{
				titulo:  "eu quero so voce",
				artista: "jorge e matheus",
				duracao: 4.2,
			},
		},
	}
	DadosPlaylist(playlist1)
}
