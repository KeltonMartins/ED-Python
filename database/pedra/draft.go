package main
import "fmt"
func main() {
    n := 0
    menorValor := 99999999
    valor := 0
    indexDoGanhador := 0
    fmt.Scanln(&n)
    
    for i:=0; i<n; i++{
        n1 := 0
        n2 := 0
        fmt.Scanf("%d %d", &n1, &n2)
        if n1 > 9 && n2 > 9{
            valor = n1 - n2
            if valor < 0 {valor *= -1}
            if menorValor > valor{
                menorValor = valor
                indexDoGanhador = i
            }
        }
    }
    fmt.Println(indexDoGanhador)
}
