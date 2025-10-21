package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func frozen(p *Pen, tam float64){
	if tam<1{
		return
	}
	for range 5{
		p.Walk(tam)
		frozen(p, tam*0.37)
		p.Walk(-tam)
		p.Rotate(360 / 5)
	}
}
func circulo(p *Pen, tam float64){
	if tam<1{
		return
	}
	p.SetDown(true)
	p.DrawCircle(tam)
	for range 6{
		p.SetDown(false)
		p.Walk(tam)
		circulo(p, tam / 3)
		p.SetDown(false)
		p.Walk(-tam)
		p.Rotate(60)
	}
}
func tree(p *Pen, tam float64){
	fator := 0.75
	ang := 35.0 + float64(randInt(0, 20))-10
	if tam < 13{
		if randInt(0, 70) < 3{
			p.SetColor(250,0,0)
			p.DrawCircle(10)
			p.SetColor(0,0,0)
		}
		return
	}
	p.SetStrokeWidth(int(tam) / 7)
	p.Walk(tam)
	p.Rotate(ang)
	tree(p, tam*fator)
	p.Rotate(-ang)
	tree(p, tam*fator)
	p.Rotate(-ang)
	tree(p, tam*fator)
	p.Rotate(ang)
	p.Walk(-tam)
}
func embua(p *Pen, tam float64) {
    if tam < 5 {
        return
    }
	p.SetStrokeWidth(max(1, int(tam) / 100))
	p.Walk(tam)
	p.Rotate(-90)
	embua(p, tam-4)
}

func quadrado(p *Pen, x, y, tam float64){
	p.SetPosition(x - tam / 2, y + tam / 2)
	for range 4{
		p.Walk(tam)
		p.Rotate(90)
	}
}
func quadrados(p *Pen, x, y, tam float64){
	if tam < 5{
		return
	}
	quadrado(p, x, y, tam)
	fator := 0.45
	quadrados(p, x - tam/2, y - tam / 2, fator * tam)
	quadrados(p, x + tam/2, y - tam / 2, fator * tam)
	quadrados(p, x - tam/2, y + tam / 2, fator * tam)
	quadrados(p, x + tam/2, y + tam / 2, fator * tam)

}

func triangulo(p *Pen, tam float64){
	if tam < 10{
		return
	}
	for range 3{
		p.Walk(tam)
		p.Rotate(-120)
		triangulo(p, tam/2)
	}
}

func trigo(p *Pen, tam float64){
	if tam < 10{
		return
	}
	fator := 0.75
	ang := 35.0
	for range 4{
		p.Walk(tam)
		p.Rotate(ang)
		trigo(p, tam*fator)
		p.Rotate(-2 * ang)
		trigo(p, tam*fator)
		p.Rotate(ang)
		p.Walk(-tam)
	}
}

func main() {
	pen := NewPen(500, 500)
	pen.SetPosition(250, 500)
	pen.SetHeading(90)
	trigo(pen, 110)
	pen.SaveToFile("trigo.svg")
	fmt.Println("SVG file created successfully.")
}
