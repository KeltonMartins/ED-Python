package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ehseguroL(matriz [][]rune, l int, num rune) bool {
	for _, celula := range matriz[l] {
		if celula == num {
			return false
		}
	}
	return true
}

func ehseguroC(matriz [][]rune, c int, num rune) bool {
	for i := range len(matriz) {
		if matriz[i][c] == num {
			return false
		}
	}
	return true
}

func ehseguroQ(matriz [][]rune, l, c int, num rune) bool {
	nl := len(matriz)
	tamanhoQ := int(math.Sqrt(float64(nl)))
	inicioLQ := l - l%tamanhoQ
	inicioCQ := c - c%tamanhoQ

	for i := range tamanhoQ {
		for j := range tamanhoQ {
			if matriz[i+inicioLQ][j+inicioCQ] == num {
				return false
			}
		}
	}
	return true
}

func teste(matriz [][]rune, l, c int, num rune) bool {
	return ehseguroC(matriz, c, num) &&
		ehseguroL(matriz, l, num) &&
		ehseguroQ(matriz, l, c, num)
}

func resolver(matriz [][]rune, index int) bool {
	nl := len(matriz)

	if index == nl*nl {
		return true
	}

	linha := index / nl
	coluna := index % nl

	if matriz[linha][coluna] != '.' {
		return resolver(matriz, index+1)
	}

	for num := 1; num <= nl; num++ {
		charNum := rune(strconv.Itoa(num)[0])

		if teste(matriz, linha, coluna, charNum) {
			matriz[linha][coluna] = charNum

			if resolver(matriz, index+1) {
				return true
			}
			matriz[linha][coluna] = '.'
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	matriz := make([][]rune, N)
	for i := range N {
		scanner.Scan()
		linha := scanner.Text()
		matriz[i] = []rune(linha)
	}

	if resolver(matriz, 0) {
		for i := range N {
			fmt.Println(string(matriz[i]))
		}
	}
}
