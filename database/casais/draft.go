package main
import "fmt"
func main() {
    var n int
    num := 0
    fmt.Scanln(&n)

    var animais []int
    var casais int

    for i := 0; i < n; i++{
        fmt.Scan(&num)
        animais = append(animais, num)
    }
    for i := 0; i<n; i++{
        for j := 1; j<n-1; j++{
            if animais[i] > 0{
                value := animais[j] * -1
                if animais[i] == value && j != i && animais[i] != 0 && animais[i] != 0{
                    animais[i] = 0
                    animais[j] = 0
                    casais++
                }
            }else{
                value := animais[i] * -1
                if value == animais[j] && j != i && animais[i] != 0 && animais[i] != 0{
                    animais[i] = 0
                    animais[j] = 0
                    casais++
                }
            }
        }
    }
    fmt.Println(casais)
}
