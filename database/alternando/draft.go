package main

import (
	"fmt"
	"slices"
)

func formatar(s []int, espadaIdx int) string {
	str := "[ "
	for i, v := range s {
		if i == espadaIdx {
			if v > 0 {
				str += fmt.Sprint(v) + "> "
			} else {
				str += "<" + fmt.Sprint(v) + " "
			}
		} else {
			str += fmt.Sprint(v) + " "
		}
	}
	str += "]"
	return str
}

func calcularIndices(s []int, cur int) (elim int, newCur int) {
	n := len(s)
	if s[cur] > 0 {
		elim = (cur + 1) % n
		cand := (elim + 1) % n
		if cand > elim {
			newCur = cand - 1
		} else {
			newCur = cand
		}
	} else {
		elim = (cur - 1 + n) % n
		cand := (elim - 1 + n) % n
		if elim > cand {
			newCur = cand
		} else {
			newCur = cand - 1
		}
	}
	return
}

func main() {
	var qtd, idxInicial, dir int
	fmt.Scan(&qtd, &idxInicial, &dir)
	guerreiros := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		guerreiros[i] = i + 1
		if dir == -1 {
			if i%2 == 0 {
				guerreiros[i] *= -1
			}
		} else {
			if i%2 != 0 {
				guerreiros[i] *= -1
			}
		}
	}
	espadaIdx := idxInicial - 1
	for len(guerreiros) > 1 {
		fmt.Println(formatar(guerreiros, espadaIdx))
		elim, novo := calcularIndices(guerreiros, espadaIdx)
		guerreiros = slices.Delete(guerreiros, elim, elim+1)
		espadaIdx = novo
	}
	fmt.Println(formatar(guerreiros, 0))
}
