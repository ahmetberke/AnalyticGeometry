package analytic_geometry

import "math"

type point struct {
	x float64
	y float64
}

type line struct {
	m float64
	point point
}

func NewPoint (x float64, y float64) point{
	p := point{x: x, y: y}
	return p
}

func NewLine (point point, m float64) line{
	l := line{m: m, point: point}
	return l
}

func (l *line) XtoY(x float64)  float64{
	y := l.m*(x - l.point.x) + l.point.y
	return y
}

func DistanceOfPoints (point1 point, point2 point) float64{
	sumx := math.Pow(point1.x + point2.x, 2)
	sumy := math.Pow(point1.y + point2.y, 2)
	distance := math.Sqrt(sumx + sumy)
	return distance
}

func DistanceOfPointToLine(point point, line line)  float64{
	d := math.Abs(line.m*point.x + point.y + line.point.y - line.m*line.point.x) / math.Sqrt(math.Pow(line.m, 2) + 1)
	return d
}

func DistanceOfLines(line1 line, line2 line)  float64{
	if line1.m != line2.m {
		return 0.0
	}
	d := math.Abs(line1.point.y - line1.m*line1.point.x - line2.point.y + line2.m*line2.point.x) / math.Sqrt(math.Pow(line1.m, 2) + 1)
	return d
}