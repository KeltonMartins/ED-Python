package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func MagicSearch(slice []int, value int) int {
	if len(slice) == 0 {
		return 0
	}
	inicio := 0
	fim := len(slice) - 1
	indexDoUltimo := 0
	existe := false
	for inicio <= fim {
		meio := (inicio + fim) / 2

		if slice[meio] == value {
			indexDoUltimo = meio
			inicio = meio + 1 
			existe = true
		} else if slice[meio] < value {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	if !existe {
		for i:= range slice{
			if value > slice[i]{
				indexDoUltimo = i + 1
			}
		}
	}
	return indexDoUltimo
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
