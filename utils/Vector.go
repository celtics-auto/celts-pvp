package utils

type Vector struct {
	X int
	Y int
}

type BoundingBox struct {
	V0 Vector // Top-Left vertice
	V1 Vector // Bottom-Right vertice
}

func NewVector(x, y int) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

func NewBoundigBox(vTopLeft, vBotRight Vector) *BoundingBox {
	return &BoundingBox{
		V0: vTopLeft,
		V1: vBotRight,
	}
}

func CheckBoxCollision(box1, box2 *BoundingBox) bool {
	return box1.V0.X <= box2.V1.X &&
		box1.V1.X >= box2.V0.X &&
		box1.V0.Y <= box2.V1.Y &&
		box1.V1.Y >= box2.V0.Y
}
