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

func estafora(grid [][]byte, p Pos, letra byte, word string) bool {
	nl := len(grid)
	nc := len(grid[0])

	return (p.l < 0 || p.l >= nl || p.c < 0 || p.c >= nc || grid[p.l][p.c] != letra || len(word) < 1)
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode

func procura(grid [][]byte, p Pos, word string) int {
	if len(word) == 0 {
		return 0
	}
	if estafora(grid, p, word[0], word) {
		return 1
	}

	letraOriginal := grid[p.l][p.c]
	grid[p.l][p.c] = '.'

	for _, viz := range vizinhos(p) {
		if procura(grid, viz, word[1:]) == 0 {
			grid[p.l][p.c] = letraOriginal
			return 0
		}
	}

	grid[p.l][p.c] = letraOriginal
	return 1
}

func exist(grid [][]byte, word string) bool {
	nl := len(grid)
	nc := len(grid[0])

	for l := range nl {
		for c := range nc {
			if procura(grid, Pos{l: l, c: c}, word) < 1 {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
