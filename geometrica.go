package figuras

type figurasGeometricas interface {
	area() float64
	perimetro() float64
}

func CalcularArea(figura figurasGeometricas) float64 {
	return figura.area()
}

func CalcularPerimetro(figura figurasGeometricas) float64 {
	return figura.perimetro()
}
