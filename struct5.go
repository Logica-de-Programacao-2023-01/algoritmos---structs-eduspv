package main

import "fmt"

//Usando as mesmas structs do exercício anterior,
//escreva uma função que receba um título de música como parâmetro e retorne uma lista das Playlists
//que possuem esse título.

type playlists struct {
	nome    string
	musicas []struct {
		titulo  string
		artista string
		duracao float64
	}
}

func print(playlists1 playlists, musica string) string {
	for _, songs := range playlists1.musicas {
		if songs.titulo == musica {
			return playlists1.nome
		}

	}
	return ""
}

func main() {
	musica := ""
	fmt.Println("digite uma musica que voce deseja verificar em quais playlists ela existe")
	fmt.Scan(&musica)
	playlist1 := playlists{
		nome: "pl1",
		musicas: []struct {
			titulo  string
			artista string
			duracao float64
		}{{
			titulo:  "vasco",
			artista: "vasco",
			duracao: 4.5,
		},
		},
	}
	playlist2 := playlists{
		nome: "pl2",
		musicas: []struct {
			titulo  string
			artista string
			duracao float64
		}{{
			titulo:  "da gama",
			artista: "Vamos perder",
			duracao: 5.6,
		},
		},
	}

	resultado1 := print(playlist1, musica)
	resultado2 := print(playlist2, musica)

	if resultado1 != "" {
		fmt.Printf("A música \"%s\" está na playlist %s\n", musica, resultado1)
	} else {
		fmt.Printf("A música \"%s\" não está na playlist %s\n", musica, playlist1.nome)
	}

	if resultado2 != "" {
		fmt.Printf("A música \"%s\" está na playlist %s\n", musica, resultado2)
	} else {
		fmt.Printf("A música \"%s\" não está na playlist %s\n", musica, playlist2.nome)
	}
}
