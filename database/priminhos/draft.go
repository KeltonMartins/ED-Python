package main

import "fmt"

func Primo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	qtdPrimos := 0
	fmt.Scan(&qtdPrimos)
	primo := 2
	var primos []int
	for len(primos) < qtdPrimos {
		if Primo(primo) {
			primos = append(primos, primo)
		}
		primo++
	}

	saida := "["
	for i := range primos {
		if i != 0 {
			saida += ", " + fmt.Sprint(primos[i])
			continue
		}
		saida += fmt.Sprint(primos[i])
	}
	saida += "]"
	fmt.Println(saida)
}
