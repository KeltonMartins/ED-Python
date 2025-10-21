package main
import "fmt"
func main() {
    var helicoptero, policial, fugitivo, direcao int
    fmt.Scanln(&helicoptero, &policial, &fugitivo, &direcao)

    distP := 0
    distF := 0
    
    if direcao < 1 {
        for true {
            if policial != helicoptero{
                distP++
                policial--
                if(policial < 1){
                    policial = 15
                }
            }
            if fugitivo != helicoptero{
                distF++
                fugitivo--
                if(fugitivo < 1){
                    fugitivo = 15
                }
            }
            if fugitivo == helicoptero || policial == helicoptero{
                break;
            }
        }
    }
    if direcao == 1{
        for true {
            if policial != helicoptero{
                distP++
                policial++
                if(policial > 15){
                    policial = 1
                }
            }
            if fugitivo != helicoptero{
                distF++
                fugitivo++
                if(fugitivo > 15){
                    fugitivo = 1
                }
            }
            if fugitivo == helicoptero || policial == helicoptero{
                break;
            }
        }
    }
    if policial == helicoptero{
        fmt.Println("N")
    }else{
        fmt.Println("S")
    }
}
