package main
import "fmt"

func ehPrimo(num int)bool{
    if num < 2{
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func pegarNumeroPrimo(n, num int) int{
    if n == 0{
        return num -1
    }
    if ehPrimo(num){
        return pegarNumeroPrimo(n-1, num+1)
    }
    return pegarNumeroPrimo(n, num+1)
}

func main() {
    n:=0
    fmt.Scan(&n)
    fmt.Println(pegarNumeroPrimo(n, 2))
}
