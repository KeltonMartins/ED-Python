package main
import "fmt"
func DivAoContrário(num int) string{
    if num == 0{
        return ""
    }
    div := num/2
    mod := num%2
    return DivAoContrário(div) + fmt.Sprintln(div, mod)
}

func main() {
    num := 0
    fmt.Scan(&num)
    fmt.Print(DivAoContrário(num))
}
