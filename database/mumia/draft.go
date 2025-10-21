package main
import "fmt"
func main() {
    var nome string
    var idade int

    fmt.Scanln(&nome)
    fmt.Scanln(&idade)

    if idade < 12{
        fmt.Println(nome, "eh crianca")
    }else if idade < 18 {
        fmt.Println(nome, "eh jovem")
    }else if idade < 65{
        fmt.Println(nome, "eh adulto")
    }else if idade < 1000{
        fmt.Println(nome, "eh idoso")
    }else{
        fmt.Println(nome, "eh mumia")
    }
    
}
