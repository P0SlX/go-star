package image

import (
	. "github.com/P0SLX/go-star/node"
	"image"
	"image/png"
	"os"
	"runtime"
)

type Image struct {
	File *os.File

	image.Image
}

func NewImage(path string) (*Image, error) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	return &Image{File: file, Image: img}, nil
}

func (i Image) readder(width, start, end int, nodes *[][]*Node) {

	for y := start; y < end; y++ {
		var row = make([]*Node, width)
		for x := 0; x < width; x++ {

			// Extrait les valeurs RGB du pixel
			r, g, b, _ := i.Image.At(x, y).RGBA()

			node := NewNode(x, y, NewColor(r, g, b))

			row[x] = node

		}
		(*nodes)[y] = row
	}

}

func (i Image) readderOptimized(width, height int, nodes *[][]*Node) {

	chunck := height / 10

	maxWorkers := runtime.NumCPU()

	sem := make(chan struct{}, maxWorkers)

	for j := 0; j < chunck; j++ {
		sem <- struct{}{}
		go func(width, start, end int, nodes *[][]*Node) {
			defer func() { <-sem }()
			i.readder(width, start, end, nodes)
		}(width, j*10, (j+1)*10-1, nodes)
	}

	for j := 0; j < maxWorkers; j++ {
		sem <- struct{}{}
	}

}

func (i *Image) Read() *[][]*Node {

	// Limites de l'image
	bounds := i.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var nodes = make([][]*Node, height)

	//Without go routines
	if width <= 16 && height <= 16 {
		i.readder(width, 0, height, &nodes)
	} else {
		//With go routines
		i.readderOptimized(width, height, &nodes)
	}

	return &nodes

}
