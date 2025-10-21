package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func tostrrec(vet []int) string{
	if len(vet) == 0{
		return ""
	}
	if len(vet) == 1{
		return fmt.Sprint(vet[0])
	}
	return fmt.Sprint(vet[0]) + ", " + tostrrec(vet[1:])
}
func tostrrecrev(vet []int) string{
	if len(vet) == 0{
		return ""
	}
	if len(vet) == 1{
		return fmt.Sprint(vet[0])
	}
	return tostrrecrev(vet[1:])+ ", " + fmt.Sprint(vet[0])
}

func tostr(vet []int) string {
	return "[" + tostrrec(vet) + "]"
}

func tostrrev(vet []int) string {
	return "[" + tostrrecrev(vet) + "]"
}
// reverse: inverte os elementos do slice
func rev(vet []int) []int{
	if len(vet) == 0{
		return vet
	}
	return append(rev(vet[1:]), vet[0])
}
func reverse(vet []int){
	copy(vet, rev(vet))
}
func soma(vet []int) int{
	if len(vet) == 0{
		return 0
	}
	return vet[0] + soma(vet[1:])
}
func sum(vet []int) int {
	return soma(vet)
}

func vezes(vet[]int) int{
	if len(vet) == 0{
		return 1
	}
	if len(vet) == 1{
		return vet[0]
	}
	return vet[0] * vezes(vet[1:])
}
func mult(vet []int) int {
	return vezes(vet)
}
var minor, minIdx int

func menorRec(vet []int, idx int, menor int, pos int) (int, int) {
	if idx == len(vet) {
		return menor, pos
	}
	if vet[idx] < menor {
		return menorRec(vet, idx+1, vet[idx], idx)
	}
	return menorRec(vet, idx+1, menor, pos)
}

func min(vet []int) (int, int) {
	if len(vet) == 0 {
		return -1, -1
	}
	if len(vet) == 1 {
		return 0, 0
	}
	return menorRec(vet, 1, vet[0], 0)
}


func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			_, index := min(vet)
			fmt.Println(index)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
