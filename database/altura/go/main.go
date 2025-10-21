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

// Search node with value
// Starts with node == root
// Value is the value we are looking for
// Returns the node with the value or nil if not found
func find(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}
	achado := find(node.Left, value)
	if achado != nil {
		return achado
	}

	return find(node.Right, value)
}

// node is the node we want to find the height of
// the height of a node is the number of edges on the longest path from the node to a leaf
func getHeight(node *Node) int {
	if node == nil {
		return 0
	}

	heightLeft := getHeight(node.Left)
	heightRight := getHeight(node.Right)

	if heightLeft > heightRight {
		return 1 + heightLeft
	}
	return 1 + heightRight
}

// node is the root of the tree
// level is the current level in the tree (1 for root)
// value is the value we are looking for
func calcNodeDepth(node *Node, level int, value int) int {
	if node == nil {
		return -1
	}

	if node.Value == value {
		return level
	}

	fundura := calcNodeDepth(node.Left, level+1, value)

	if fundura != -1 {
		return fundura
	}
	return calcNodeDepth(node.Right, level+1, value)
}

// --------------------------------------------------------------------
// Don't change from here
func BShow(node *Node, heranca string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, heranca+"l")
	}
	for i := 0; i < len(heranca)-1; i++ {
		if heranca[i] != heranca[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if heranca != "" {
		if heranca[len(heranca)-1] == 'l' {
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
		BShow(node.Right, heranca+"r")
	}
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

	line := scanner.Text()
	parts := strings.Split(line, " ")
	root := create(&parts)

	scanner.Scan()
	line = scanner.Text()
	fmt.Println("Arvore:")
	BShow(root, "")

	values := strings.FieldsSeq(line)
	for s := range values {
		value, _ := strconv.Atoi(s)
		node := find(root, value)
		if node != nil {
			fmt.Printf("Altura: %d, Profundidade: %d\n", getHeight(node), calcNodeDepth(root, 1, value))
		} else {
			fmt.Println("-1")
		}
	}
}
