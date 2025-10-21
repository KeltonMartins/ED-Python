package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
)

func getMen(vet []int) []int {
	var vetVolta []int
	for i := range vet{
		if vet[i] > 0 {
			vetVolta = append(vetVolta, vet[i])
		}
	}
	return vetVolta
}

func getCalmWomen(vet []int) []int {
	var vetVolta []int
	for i := range vet{
		if vet[i] > -10 &&  vet[i] < 0{
			vetVolta = append(vetVolta, vet[i])
		}
	}
	return vetVolta
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {
	var vet2 []int
	for i := range vet{
		if vet[i] < 0{
			vet2 = append(vet2, vet[i]*-1)
			vet[i] = vet[i]*-1
		}
	}
	slices.Sort(vet)
	for i := range vet2{
		index := slices.Index(vet, vet2[i])
		if index != -1{
			vet[index] = vet[index]*-1
		}
	}
	return vet
}

func reverse(vet []int) []int {
	slices.Reverse(vet)
	return vet
}

func reverseInplace(vet []int) {
	_ = vet
}
func unique(vet []int) []int {
	result := []int{}
	jaViu := make(map[int]bool)
	for _, v := range vet{
		if !jaViu[v]{
			result = append(result, v)
			jaViu[v] = true
		}
	}
	return result
}

func repeated(vet []int) []int {
	result := []int{}
	repetida := make(map[int]int)
	for _, v := range vet{
		repetida[v]++
		if repetida[v] > 1{
			result = append(result, v)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			printVec(reverse(str2vet(args[1])))
		case "reverse_inplace":
			vet := str2vet(args[1])
			reverseInplace(vet)
			printVec(vet)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

