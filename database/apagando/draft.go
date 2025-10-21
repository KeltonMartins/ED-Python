package main
import(
    "fmt"
    "slices"
)
func formatar [T any](slice []T, sep string)string{
    stringfinal := ""
    for i, v := range slice{
        if i != 0{
            stringfinal += sep
        }
        stringfinal += fmt.Sprint(v)
    }
    stringfinal += sep
    return stringfinal
}

func main() {
    var qtdPessoas, qtdTemQueSair int
    fmt.Scan(&qtdPessoas)
    idpessoas := make([]int, qtdPessoas)

    for i := 0; i<qtdPessoas; i++{
        fmt.Scan(&idpessoas[i])
    }

    fmt.Scan(&qtdTemQueSair)
    sair := make([]int, qtdTemQueSair)

    for i := 0; i<qtdTemQueSair; i++{
        fmt.Scan(&sair[i])
        id := slices.Index(idpessoas, sair[i])
        if id != -1{
            idpessoas = slices.Delete(idpessoas, id, id+1)
        }
    }

    fmt.Println(formatar(idpessoas, " "))
}
