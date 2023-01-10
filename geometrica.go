package figuras

type FigurasGeometricas interface {
	area() float64
	perimetro() float64
}

func CalcularArea(figura FigurasGeometricas) float64 {
	return figura.area()
}

func CalcularPerimetro(figura FigurasGeometricas) float64 {
	return figura.perimetro()
}
