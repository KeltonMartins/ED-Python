package main

import (
	"fmt"
)

const EMPTY = '.'

type Problem struct {
	data []rune
	lim  int
}

// verifica se esse valor pode ser utilizado nessa posição
func (p *Problem) fit(index int, value rune) bool {
	for i := 1; i <= p.lim; i++ {
		posicaoChecada := index - i
		if posicaoChecada < 0 {
			break
		}
		if p.data[posicaoChecada] == value {
			return false
		}
	}

	for i := 1; i <= p.lim; i++ {
		posicaoChecada := index + i
		if posicaoChecada >= len(p.data) {
			break
		}
		if p.data[posicaoChecada] == value {
			return false
		}
	}

	// cuidado para não sair dos limites, nem para o começo, nem para o fim
	return true
}

func (p *Problem) solve(index int) bool {
	// se chegou ao fim, retorne true
	if index >= len(p.data) {
		return true
	}
	// se não é EMPTY, vá para o próximo
	if p.data[index] != EMPTY {
		return p.solve(index + 1)
	}

	// se é EMPTY
	//    faça um laço e chame a recursão para cada valor possível
	for i := p.lim; i >= 0; i-- {
		r := rune('0' + i)

		if p.fit(index, r) {
			p.data[index] = r

			if p.solve(index + 1) {
				//    se algum der certo, então retorne verdade
				return true
			}

			// se nenhum deu certo, recoloque vazio e retorne falso
			p.data[index] = EMPTY
		}
	}
	return false
}

func main() {
	var input string
	var lim int
	fmt.Scan(&input, &lim)
	prob := Problem{data: []rune(input), lim: lim}
	prob.solve(0)
	fmt.Println(string(prob.data))
}
