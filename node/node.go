package node

type Node struct {
	X int
	Y int

	Color *Color

	G, H, F   float64
	Neighbors []*Node
	Parent    *Node
	IsWall    bool
}

func NewNode(x, y int, color *Color) *Node {
	return &Node{
		X:      x,
		Y:      y,
		Color:  color,
		IsWall: color.R == 0 && color.G == 0 && color.B == 0,
	}
}
