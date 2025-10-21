package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func soma(a *Node) int {
	if a == nil {
		return 0
	}

	return a.Value + soma(a.Left) + soma(a.Right)
}

func menor(a *Node) int {
	if a == nil {
		return math.MaxInt32
	}

	menorEsquerda := menor(a.Left)
	menorDireita := menor(a.Right)

	menorValor := a.Value
	if menorEsquerda < menorValor {
		menorValor = menorEsquerda
	}

	if menorDireita < menorValor {
		menorValor = menorDireita
	}

	return menorValor
}

func create(parts *[]string) *Node {
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	a := create(&parts)

	somaTotal := soma(a)
	menorValor := menor(a)

	fmt.Println(somaTotal, menorValor)
}
