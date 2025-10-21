package main
import "fmt"

func backtracking(vet []int, index, alvo, soma int) bool{
    if soma == alvo{
        return true
    }
    if soma > alvo || index >= len(vet){
        return false
    }
    if backtracking(vet,index+1, alvo, soma+vet[index]) ||
     backtracking(vet, index+1, alvo, soma){
        return true
    }
    return false
}

func main() {
    var qtd, alvo int
    fmt.Scan(&qtd, &alvo)
    vet := make([]int, qtd)
    for i := range qtd{
        fmt.Scan(&vet[i])
    }
    if(backtracking(vet, 0, alvo, 0)){
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }
}
