package image

import (
	. "github.com/P0SLX/go-star/node"
	"image"
	"image/color"
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

	rest := height % 10

	maxWorkers := runtime.NumCPU()

	sem := make(chan struct{}, maxWorkers)

	for j := 0; j < chunck; j++ {
		sem <- struct{}{}
		go func(width, start, end int, nodes *[][]*Node) {
			defer func() { <-sem }()
			i.readder(width, start, end, nodes)
		}(width, j*10, (j+1)*10, nodes)
	}

	for j := 0; j < maxWorkers; j++ {
		sem <- struct{}{}
	}

	if rest > 0 {
		i.readder(width, chunck*10, chunck*10+rest, nodes)
	}

}

func (i *Image) findNeighbors(nodes [][]*Node) {

	for y := range nodes {
		//TODO : Faire un tableau avec make
		for x := range nodes[y] {
			// Haut
			if y > 0 {
				nodes[y][x].Neighbors = append(nodes[y][x].Neighbors, nodes[y-1][x])
			}

			// Bas
			if y < len(nodes)-1 {
				nodes[y][x].Neighbors = append(nodes[y][x].Neighbors, nodes[y+1][x])
			}

			// Gauche
			if x > 0 {
				nodes[y][x].Neighbors = append(nodes[y][x].Neighbors, nodes[y][x-1])
			}

			// Droite
			if x < len(nodes[y])-1 {
				nodes[y][x].Neighbors = append(nodes[y][x].Neighbors, nodes[y][x+1])
			}
		}
	}
}

func nodeToImage(nodes [][]*Node) image.Image {
	width, height := len(nodes), len(nodes[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := range nodes {
		for j := range nodes[i] {
			img.Set(j, i, color.RGBA{
				R: nodes[i][j].Color.R,
				G: nodes[i][j].Color.G,
				B: nodes[i][j].Color.B,
				A: 255,
			})
		}
	}
	return img
}

// FindStartAndEndNode Détecte le point de départ et d'arrivée
//
// Le point de départ est un pixel vert (0, 255, 0)
// Le point d'arrivée est un pixel rouge (255, 0, 0)
// On boucle sur chaque node, et on renvoie le pointeur
// du premier pixel vert et du premier pixel rouge
func (i Image) FindStartAndEndNode(nodes [][]*Node) (*Node, *Node) {
	var start *Node
	var end *Node

	for y := range nodes {
		for x := range nodes[y] {
			// Point de départ
			if nodes[y][x].Color.IsStartPoint() {
				start = nodes[y][x]
			}

			// Point d'arrivée
			if nodes[y][x].Color.IsEndPoint() {
				end = nodes[y][x]
			}

			// Tout trouvé ? Pas besoin de continuer
			if start != nil && end != nil {
				return start, end
			}
		}
	}

	return nil, nil
}

// Read Extrait les pixels d'une image en un tableau 2D de Node
//
// La fonction décode l'image, obtient les limites,
// initialise un tableau 2D de Node, et boucle sur chaque pixel.
func (i *Image) Read() [][]*Node {

	// Limites de l'image
	bounds := i.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var nodes = make([][]*Node, height)
	defer i.findNeighbors(nodes)

	//Without go routines
	if width <= 16 && height <= 16 {
		i.readder(width, 0, height, &nodes)
	} else {
		//With go routines
		i.readderOptimized(width, height, &nodes)
	}

	return nodes

}

func (i Image) Save(nodes [][]*Node, filename string) error {
	img := nodeToImage(nodes)

	out, err := os.Create("./ressources/" + filename)

	if err != nil {
		return err
	}

	defer out.Close()

	err = png.Encode(out, img)

	return err
}
