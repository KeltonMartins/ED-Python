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
		{p.l + 1, p.c},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l, p.c - 1},
	}
}
func estaFora(grid [][]byte, p Pos) bool {
	nl := len(grid)
	nc := len(grid[0])
	return (p.l < 0 || p.l >= nl || p.c < 0 || p.c >= nc || grid[p.l][p.c] == '0')
}
func numIslands(grid [][]byte) int {
	nl := len(grid)
	nc := len(grid[0])
	contailhas := 0
	for l := range nl {
		for c := range nc {
			if grid[l][c] == '1' {
				contailhas++
				afundar(grid, Pos{l: l, c: c})
			}
		}
	}
	return contailhas
}

func afundar(grid [][]byte, p Pos) {
	if estaFora(grid, p) {
		return
	}
	grid[p.l][p.c] = '0'

	for _, viz := range vizinhos(p) {
		afundar(grid, viz)
	}
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
