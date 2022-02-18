package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type side int

const (
	SidesTriangle side = 3
	SidesSquare   side = 4
	SidesCircle   side = 0
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type triangle float64
type square float64
type circle float64

type squarer interface {
	area() float64
}

func (t triangle) area() float64 {
	var res, l float64
	l = float64(t)
	res = math.Sqrt(3) / 4 * math.Pow(l, 2)
	return res
}

func (c circle) area() float64 {
	var res, r float64
	r = float64(c)
	res = math.Pi * math.Pow(r, 2)
	return res
}

func (s square) area() float64 {
	var res, a float64
	a = float64(s)
	res = math.Pow(a, 2)
	return res
}

func (s side) buildSquarer(sideLen float64) squarer {
	if s == SidesTriangle {
		return triangle(sideLen)
	} else if s == SidesSquare {
		return square(sideLen)
	} else if s == SidesCircle {
		return circle(sideLen)
	} else {
		return nil
	}
}

func CalcSquare(sideLen float64, sidesNum side) float64 {
	var sq squarer
	sq = sidesNum.buildSquarer(sideLen)
	if sq == nil {
		return 0
	} else {
		return sq.area()
	}
}
