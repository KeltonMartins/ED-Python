package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func vizinhos(p Pos) []Pos {
	return []Pos{
		{p.l, p.c - 1},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l + 1, p.c},
	}
}

func estaFora(mat [][]rune, p Pos) bool {
	nl := len(mat)
	nc := len(mat[0])
	return (p.l < 0 || p.l >= nl || p.c < 0 || p.c >= nc || mat[p.l][p.c] != '.')
}

func achar(mat [][]rune, inicio, fim Pos) {
	nl := len(mat)
	nc := len(mat[0])

	vis := make([][]bool, nl)
	for i := range vis {
		vis[i] = make([]bool, nc)
	}

	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()
	caminho.Push(inicio)

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		vis[atual.l][atual.c] = true

		if atual == fim {
			break
		}

		var validos []Pos
		for _, viz := range vizinhos(atual) {
			if mat[viz.l][viz.c] != '#' && !vis[viz.l][viz.c] {
				validos = append(validos, viz)
			}
		}

		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			becos.Push(caminho.Pop())
		}
	}
	for _, p := range caminho.data {
		if mat[p.l][p.c] == ' ' || mat[p.l][p.c] == 'I' || mat[p.l][p.c] == 'F' {
			mat[p.l][p.c] = '.'
		}
	}

	for _, p := range becos.data {
		if mat[p.l][p.c] == '.' {
			mat[p.l][p.c] = ' '
		}
	}
	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			fmt.Print(string(mat[i][j]))
		}
		fmt.Println()
	}
}

func main() {
	nl, nc := 0, 0
	fmt.Scan(&nl, &nc)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	mat := make([][]rune, 0, nl)

	for i := 0; i < nl; i++ {
		scanner.Scan()
		linha := []rune(scanner.Text())
		mat = append(mat, linha)
	}
	var inicio, fim Pos

	for i := range nl {
		for j := range nc {
			if mat[i][j] == 'I' {
				inicio = Pos{l: i, c: j}
				mat[i][j] = ' '
			}
			if mat[i][j] == 'F' {
				fim = Pos{l: i, c: j}
				mat[i][j] = ' '
			}
		}
	}
	achar(mat, inicio, fim)
}