package main
import "fmt"
type Gomo struct{
    x int
    y int
}

func main() {
    var Q, num1, num2 int
    D := ""
    fmt.Scan(&Q, &D)

    var cobra []Gomo
    for i := 0; i<Q; i++{
        fmt.Scan(&num1, &num2)
        cobra = append(cobra, Gomo{num1, num2})
    }
    for i:= Q-1; i>0; i--{
        cobra[i] = cobra[i-1]
    }

    if D == "L"{
        cobra[0].x -= 1
    }else if D == "U"{
        cobra[0].y -= 1
    }else if D == "R"{
        cobra[0].x += 1
    }else if D == "D"{
        cobra[0].y += 1
    }

    for i:= 0; i<Q; i++{
        fmt.Println(cobra[i].x, cobra[i].y)
    }
}
