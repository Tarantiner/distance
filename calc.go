package main

import (
	"math"
)

const R = 6371.009

type P struct {
	X float64
	Y float64
}

func Raid(d float64) float64 {
	return d * math.Pi / 180.0
}

// 大圆坐标下，两点距离
func Distance(p0, p1 P) float64 {
	latA := p0.Y
	lngA := p0.X
	latB := p1.Y
	lngB := p1.X

	phi1 := Raid(latA)
	lambda1 := Raid(lngA)
	phi2 := Raid(latB)
	lambda2 := Raid(lngB)

	deltaLambda := math.Abs(lambda1 - lambda2)

	centralAngle := math.Atan2(
		math.Sqrt(math.Pow(math.Cos(phi2) * math.Sin(deltaLambda), 2.0) + math.Pow(math.Cos(phi1) * math.Sin(phi2) - math.Sin(phi1) * math.Cos(phi2) * math.Cos(deltaLambda), 2.0)),
		math.Sin(phi1) * math.Sin(phi2) + math.Cos(phi1) * math.Cos(phi2) * math.Cos(deltaLambda))
	return (R * centralAngle) * 1000
}
