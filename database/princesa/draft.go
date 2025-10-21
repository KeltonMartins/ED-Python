package main

import (
	"fmt"
	"slices"
)

func formatar(slice []int, k int, sep string) string {
	stringvolta := "[ "
	for _, v := range slice {
		if v == k {
			stringvolta += fmt.Sprint(v) + "> "
		} else {
			stringvolta += fmt.Sprint(v) + sep
		}
	}
	stringvolta += "]"
	return stringvolta
}

func main() {
	var n, matador int
	fmt.Scanf("%d %d", &n, &matador)

	var matar []int
	for i := 1; i <= n; i++ {
		matar = append(matar, i)
	}
	matador -= 1
	fmt.Println(formatar(matar, matar[matador], " "))

	for len(matar) > 1 {
		indexDaMorte := (matador + 1) % len(matar)
		matar = slices.Delete(matar, indexDaMorte, indexDaMorte+1)
		matador = indexDaMorte % len(matar)
		fmt.Println(formatar(matar, matar[matador], " "))
	}
}
