package point

import "math"

type Point struct {
	x, y float32
}

func New(x, y float32) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Вычисляет расстояние между точками. Чистейшая школьная геометрия.
func (p *Point) Distance(point *Point) float32 {
	return float32(math.Sqrt(math.Pow(float64(p.y-point.y), 2) + math.Pow(float64(p.x-point.x), 2)))
}
