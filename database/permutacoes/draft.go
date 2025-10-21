package main
import (
    "fmt"
    "sort"
)
func permutar(s []byte, inicio int, result *[]string){
    if inicio == len(s){
        *result = append(*result, string(s))
        return
    }
    for i:= inicio; i<len(s); i++{
        s[i], s[inicio] = s[inicio], s[i]
        permutar(s, inicio+1, result)
        s[i], s[inicio] = s[inicio], s[i]
    }
}

func main() {
    var s string
    var result []string
    fmt.Scan(&s)
    permutar([]byte(s), 0, &result)
    sort.Strings(result)

    for _, r := range result{
        fmt.Println(r)
    }
    
}