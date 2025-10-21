package main
import "fmt"

func Join[T any](slice []T, sep string) string {
    result := ""
    if len(slice) < 1{
        return "N"
    }
    for i, v := range slice{
        if i != 0 {
            result += sep
        }
        result += fmt.Sprint(v)
    }
    return result
}

func main() {
    var qtdTotais, qtdBaruel, num int
    fmt.Scan(&qtdTotais, &qtdBaruel)

    var numFigs []int
    var falta []int
    var repetidas []int
    possui := make(map[int]int)

    for i := 0; i < qtdBaruel; i++{
        fmt.Scan(&num)
        possui[num]++
        numFigs = append(numFigs, num)
    }

    for i := 0; i < qtdBaruel-1; i++{
        if numFigs[i] == numFigs[i+1]{
            repetidas = append(repetidas, numFigs[i])
        }
    }

    for i := 1; i <= qtdTotais; i++{
        if possui[i] == 0{
            falta = append(falta, i)
        }
    }

    fmt.Println(Join(repetidas, " "))
    fmt.Println(Join(falta, " "))
}