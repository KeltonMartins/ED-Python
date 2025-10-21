package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x <= 1{
		return false
	}
	if div >= x{
		return true
	}
	if x % div == 0 {
		return false
	}
	div += 1
	return eh_primo(x, div)
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
