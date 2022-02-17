package soluton

import "math"

type intSidesNumType int

const SidesTriangle intSidesNumType = 3
const SidesSquare intSidesNumType = 4
const SidesCircle intSidesNumType = 0

func CalcSquare(sideLen float64, sidesCount intSidesNumType) float64 {
	if sidesCount == SidesTriangle {
		return triangleSquare(sideLen)
	}
	if sidesCount == SidesSquare {
		return squareSquare(sideLen)
	}
	if sidesCount == SidesCircle {
		return circleSquare(sideLen)
	}

	return 0
}

func triangleSquare(sideLen float64) float64 {
	p := (sideLen * float64(SidesTriangle)) / 2
	return math.Sqrt(p * math.Pow(p-sideLen, 3))
}

func squareSquare(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func circleSquare(sideLen float64) float64 {
	return math.Pi * math.Pow(sideLen, 2)
}
