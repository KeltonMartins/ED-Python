package main
import "fmt"

func rodar(vet []int, r int) []int {
    n := len(vet)
    if n == 0 || r <= 0 {
        return vet
    }
    r = r % n
    return append(vet[n-r:], vet[:n-r]...)
}
func printar (vet []int)string{
    saida := "["
    for i:= range vet{
        saida += " " + fmt.Sprint(vet[i])
    }
    saida += " ]"
    return saida
}

func main() {
    var n, r int
    fmt.Scan(&n)
    fmt.Scan(&r)
    vet := make([]int, n)
    for i:= range n{
        fmt.Scan(&vet[i])
    }
    fmt.Println(printar(rodar(vet, r)))
}
