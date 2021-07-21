package utils

type Vector struct {
	X int
	Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}
