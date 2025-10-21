package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	resultado := []Pair{}
	repete := make(map[int]int)
	for _, v := range vet {
		if v < 0 {
			v *= -1
		}
		repete[v]++
	}
	for i, v := range repete {
		resultado = append(resultado, Pair{One: i, Two: v})
	}
	slices.SortFunc(resultado, func(a, b Pair) int {
		return cmp.Compare(a.One, b.One)
	})
	return resultado
}

func teams(vet []int) []Pair {
	resultado := []Pair{}
	if len(vet) == 0 {
		return resultado
	}

	atual := vet[0]
	contador := 1

	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			contador++
		} else {
			resultado = append(resultado, Pair{One: atual, Two: contador})
			atual = vet[i]
			contador = 1
		}
	}
	resultado = append(resultado, Pair{One: atual, Two: contador})
	return resultado
}

func mnext(vet []int) []int {
	doLado := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if (i > 0 && vet[i-1] < 0) || (i < len(vet)-1 && vet[i+1] < 0) {
				doLado[i] = 1
			}
		}
	}
	return doLado
}

func alone(vet []int) []int {
	doLado := make([]int, len(vet))
	for i, v := range vet {
		if v > 0 {
			temMulher := false

			if i > 0 && vet[i-1] < 0 {
				temMulher = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				temMulher = true
			}
			if !temMulher {
				doLado[i] = 1
			}
		}
	}
	return doLado
}

func couple(vet []int) int {
	men := make(map[int]int)
	women := make(map[int]int)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for _, stress := range vet {
		level := abs(stress)
		if stress > 0 {
			men[level]++
		} else if stress < 0 {
			women[level]++
		}
	}
	couples := 0
	for level, countMen := range men {
		if countWomen, exists := women[level]; exists {
			if countMen < countWomen {
				couples += countMen
			} else {
				couples += countWomen
			}
		}
	}
	return couples
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	n, m := len(vet), len(seq)
	if m == 0 {
		return 0
	}
	for i := 0; i <= n-m; i++ {
		found := true
		for j := 0; j < m; j++ {
			if vet[i+j] != seq[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	rm := make(map[int]bool)
	for _, pos := range posList {
		rm[pos] = true
	}
	res := []int{}
	for i, v := range vet {
		if !rm[i] {
			res = append(res, v)
		}
	}
	return res
}


func clear(vet []int, value int) []int {
	res := []int{}
	for _, v := range vet {
		if v != value {
			res = append(res, v)
		}
	}
	return res
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}