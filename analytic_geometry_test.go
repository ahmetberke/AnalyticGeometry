package analytic_geometry

import "testing"

func TestNewPoint(t *testing.T) {
	x := 15.0
	y := 23.1
	point := NewPoint(x,y)
	if point.x != x || point.y != y {
		t.Errorf("newPoint(%f,%f) = point{x=%f,y=%f};", x, y, point.x, point.y)
	}
}

func TestNewLine(t *testing.T) {
	point := NewPoint(10,20)
	m := 1.0
	line := NewLine(point, m)
	if line.m != m || line.point != point {
		t.Errorf("m = %f, point ( x=%f, y=%f )", line.m, line.point.x, line.point.y)
	}
}

func TestLine_XtoY(t *testing.T) {
	point := NewPoint(10,20)
	m := 1.0
	line := NewLine(point, m)
	x := 20.0
	y := line.XtoY(x)
	if (y != 30.0) {
		t.Errorf("y = %f - must be 30", y)
	}
}

func TestDistanceOfPoints(t *testing.T) {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3,4)
	distance := DistanceOfPoints(p1, p2)
	if distance != 5 {
		t.Errorf("distance must be 5")
	}
}

func TestDistanceOfPointToLine(t *testing.T) {
	p := NewPoint(0,0)
	l := NewLine(p, 1)
	d := DistanceOfPointToLine(p, l)
	if d != 0.0 {
		t.Errorf("distance = %f, must be 0", d)
	}
}

func TestDistanceOfLines(t *testing.T) {
	p1 := NewPoint(0,0)
	l1 := NewLine(p1, 1)
	l2 := NewLine(p1, 2)

	d := DistanceOfLines(l1, l2)
	if d != 0.0 {
		t.Errorf("distance = %f, must be 0", d)
	}

	l2 = NewLine(p1, 1)
	d = DistanceOfLines(l1, l2)
	if d != 0.0 {
		t.Errorf("distance = %f, must be 5", d)
	}
}

