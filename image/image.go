package image

import (
	. "github.com/P0SLX/go-star/node"
	"github.com/P0SLX/go-star/utils"
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

func (i Image) Reader(width, start, end int, nodes *[][]*Node) {

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

func (i Image) ReaderOptimized(width, height int, nodes *[][]*Node) {

	chunk, rest, factor := utils.FindBestChunck(height)

	maxWorkers := runtime.NumCPU()

	sem := make(chan struct{}, maxWorkers)

	for j := 0; j < chunk; j++ {
		sem <- struct{}{}
		go func(width, start, end int, nodes *[][]*Node) {
			defer func() { <-sem }()
			i.Reader(width, start, end, nodes)
		}(width, j*factor, (j+1)*factor, nodes)
	}

	for j := 0; j < maxWorkers; j++ {
		sem <- struct{}{}
	}

	if rest > 0 {
		i.Reader(width, chunk*factor, chunk*factor+rest, nodes)
	}

}

func (i *Image) findNeighbors(nodes [][]*Node) {
	for y := range nodes {
		for x := range nodes[y] {
			// Haut
			if y > 0 {
				nodes[y][x].Neighbors[0] = nodes[y-1][x]
			}

			// Bas
			if y < len(nodes)-1 {
				nodes[y][x].Neighbors[1] = nodes[y+1][x]
			}

			// Gauche
			if x > 0 {
				nodes[y][x].Neighbors[2] = nodes[y][x-1]
			}

			// Droite
			if x < len(nodes[y])-1 {
				nodes[y][x].Neighbors[3] = nodes[y][x+1]
			}
		}
	}
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
		i.Reader(width, 0, height, &nodes)
	} else {
		//With go routines
		i.ReaderOptimized(width, height, &nodes)
	}

	return nodes

}

func (i Image) Save(nodes [][]*Node, filename string) error {
	img := utils.NodeToImage(nodes)

	out, err := os.Create("./ressources/" + filename)

	if err != nil {
		return err
	}

	defer out.Close()

	err = png.Encode(out, img)

	return err
}
