package main

import (
	"fmt"
	"os"
)

type SVG struct {
	width, height int
	elements      []string
}

func NewSVG(width, height int) *SVG {
	return &SVG{
		width:    width,
		height:   height,
		elements: []string{},
	}
}

func (s *SVG) DesenharLinha(x1, y1, x2, y2 int, stroke string, strokeWidth int) {
	s.elements = append(s.elements, fmt.Sprintf(
		"<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" stroke=\"%s\" stroke-width=\"%d\" />",
		x1, y1, x2, y2, stroke, strokeWidth))
}

func (s *SVG) DesenharEllipse(cx, cy, rx, ry int, stroke, fill string) {
	s.elements = append(s.elements, fmt.Sprintf(
		"<ellipse cx=\"%d\" cy=\"%d\" rx=\"%d\" ry=\"%d\" stroke=\"%s\" fill=\"%s\" />",
		cx, cy, rx, ry, stroke, fill))
}

func (s *SVG) DesenharRect(x, y, width, height int, stroke, fill string) {
	s.elements = append(s.elements, fmt.Sprintf(
		"<rect x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" stroke=\"%s\" fill=\"%s\" />",
		x, y, width, height, stroke, fill))
}

func (s *SVG) SalvarArquivo(nomeArquivo string) error {
	file, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">\n", s.width, s.height))
	if err != nil {
		return err
	}

	for _, element := range s.elements {
		_, err = file.WriteString(element + "\n")
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString("</svg>")
	return err
}
