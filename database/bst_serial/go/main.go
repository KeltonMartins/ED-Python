package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func colocarNo(a *Node, value int) *Node {
	if a == nil {
		return &Node{Value: value}
	}

	if value < a.Value {
		a.Left = colocarNo(a.Left, value)
	} else if value > a.Value {
		a.Right = colocarNo(a.Right, value)
	}

	return a
}

func BstInsert(values []int) *Node {
	var avre *Node = nil

	for _, value := range values {
		avre = colocarNo(avre, value)
	}

	return avre
}

// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial
func Serialize(root *Node) string {
	var partes []string
	serializar(root, &partes)
	return strings.Join(partes, " ")
}

func serializar(a *Node, partes *[]string) {
	if a == nil {
		*partes = append(*partes, "#")
		return
	}

	*partes = append(*partes, strconv.Itoa(a.Value))
	serializar(a.Left, partes)
	serializar(a.Right, partes)
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}
