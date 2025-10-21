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

func getNeig(p Pos) []Pos {
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
	return (p.l < 0 || p.l >= nl || p.c < 0 || p.c >= nc || mat[p.l][p.c] != '#')
}

func tocarFogo(mat [][]rune, p Pos) {
	if estaFora(mat, p) {
		return
	}
	mat[p.l][p.c] = 'o'
	for _, viz := range getNeig(p) {
		tocarFogo(mat, viz)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	mat := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		linha := []rune(scanner.Text())
		mat = append(mat, linha)
	}
	tocarFogo(mat, Pos{l: lfire, c: cfire})
	showMat(mat)
}

func showMat(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
