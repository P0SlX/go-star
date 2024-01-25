package node

import (
	"github.com/P0SLX/go-star/utils"
	"image"
	"image/color"
	"math"
	"runtime"
)

var (
	ALPHA uint8 = 255
)

type Node struct {
	X int
	Y int

	Color *Color

	G, H, F float64

	Neighbors []*Node
	Parent    *Node
	IsWall    bool
	Already   bool // Déjà visité
}

func NewNode(x, y int, color *Color) *Node {
	return &Node{
		X:         x,
		Y:         y,
		Color:     color,
		IsWall:    color.R == 0 && color.G == 0 && color.B == 0,
		Neighbors: make([]*Node, 4),
	}
}

// Heuristic calcule la distance Euclidienne entre 2 points
func (n Node) Heuristic(dest *Node) float64 {
	xSquare := float64(n.X-dest.X) * float64(n.X-dest.X)
	ySquare := float64(n.Y-dest.Y) * float64(n.Y-dest.Y)
	return math.Sqrt(xSquare + ySquare)
}

type Nodes [][]*Node

func (n Nodes) ToImageOptimized() image.Image {
	width, height := len(n), len(n[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	chunk, rest, factor := utils.FindBestChunck(height)

	maxWorkers := runtime.NumCPU()

	sem := make(chan struct{}, maxWorkers)

	for j := 0; j < chunk; j++ {
		sem <- struct{}{}
		go func(width, start, end int, nodes [][]*Node) {
			defer func() { <-sem }()
			for y := start; y < end; y++ {
				for x := 0; x < width; x++ {
					img.Set(x, y, color.RGBA{
						R: nodes[y][x].Color.R,
						G: nodes[y][x].Color.G,
						B: nodes[y][x].Color.B,
						A: ALPHA,
					})
				}
			}

		}(width, j*factor, (j+1)*factor, n)
	}

	for j := 0; j < maxWorkers; j++ {
		sem <- struct{}{}
	}

	if rest > 0 {
		for y := chunk * factor; y < chunk*factor+rest; y++ {
			for x := 0; x < width; x++ {
				img.Set(x, y, color.RGBA{
					R: n[y][x].Color.R,
					G: n[y][x].Color.G,
					B: n[y][x].Color.B,
					A: ALPHA,
				})
			}
		}
	}

	return img
}

func (n Nodes) ToImage() image.Image {
	width, height := len(n), len(n[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range n {
		for x := range n[y] {
			img.Set(x, y, color.RGBA{
				R: n[y][x].Color.R,
				G: n[y][x].Color.G,
				B: n[y][x].Color.B,
				A: ALPHA,
			})
		}
	}
	return img
}
