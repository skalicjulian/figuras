package figuras

type Cuadrado struct {
	Ancho float64
	Alto  float64
}

func (cu *Cuadrado) area() float64 {
	return cu.Ancho * cu.Alto
}

func (cu *Cuadrado) perimetro() float64 {
	return 2*cu.Ancho + 2*cu.Alto
}
