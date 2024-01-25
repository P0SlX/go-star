package image

import (
	. "github.com/P0SLX/go-star/node"
	"github.com/P0SLX/go-star/utils"
	"image"
	"image/png"
	"os"
	"runtime"
)

var (
	MIN_SIZE = 32
)

type Image struct {
	File *os.File

	image.Image

	Chunk  int
	Rest   int
	Factor int

	Width  int
	Height int

	Nodes Nodes
}

// NewImage Crée une nouvelle instance de Image
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

	// Limites de l'image
	bounds := img.Bounds()

	chunk, rest, factor := utils.FindBestChunck(bounds.Max.Y)

	return &Image{
		File:   file,
		Width:  bounds.Max.X,
		Height: bounds.Max.Y,
		Image:  img,
		Rest:   rest,
		Chunk:  chunk,
		Factor: factor,
		Nodes:  make([][]*Node, bounds.Max.Y),
	}, nil
}

// Reader Extrait les pixels d'une image en un tableau 2D de Node
func (i Image) Reader(start, end int) {

	for y := start; y < end; y++ {
		var row = make([]*Node, i.Width)
		for x := 0; x < i.Width; x++ {

			// Extrait les valeurs RGB du pixel
			r, g, b, _ := i.Image.At(x, y).RGBA()

			node := NewNode(x, y, NewColor(r, g, b))

			row[x] = node

		}
		i.Nodes[y] = row
	}

}

// ReaderOptimized Extrait les pixels d'une image en un tableau 2D de Node
func (i Image) ReaderOptimized() {

	maxWorkers := runtime.NumCPU()

	sem := make(chan struct{}, maxWorkers)

	for j := 0; j < i.Chunk; j++ {
		sem <- struct{}{}
		go func(start, end int) {
			defer func() { <-sem }()
			i.Reader(start, end)
		}(j*i.Factor, (j+1)*i.Factor)
	}

	for j := 0; j < maxWorkers; j++ {
		sem <- struct{}{}
	}

	if i.Rest > 0 {
		i.Reader(i.Chunk*i.Factor, i.Chunk*i.Factor+i.Rest)
	}

}

// findNeighbors Trouve les voisins de chaque node
func (i *Image) findNeighbors() {
	for y := range i.Nodes {
		for x := range i.Nodes[y] {
			// Haut
			if y > 0 {
				i.Nodes[y][x].Neighbors = append(i.Nodes[y][x].Neighbors, i.Nodes[y-1][x])
			}

			// Bas
			if y < len(i.Nodes)-1 {
				i.Nodes[y][x].Neighbors = append(i.Nodes[y][x].Neighbors, i.Nodes[y+1][x])
			}

			// Gauche
			if x > 0 {
				i.Nodes[y][x].Neighbors = append(i.Nodes[y][x].Neighbors, i.Nodes[y][x-1])
			}

			// Droite
			if x < len(i.Nodes[y])-1 {
				i.Nodes[y][x].Neighbors = append(i.Nodes[y][x].Neighbors, i.Nodes[y][x+1])
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
func (i Image) FindStartAndEndNode() (*Node, *Node) {
	var start *Node
	var end *Node

	for y := range i.Nodes {
		for x := range i.Nodes[y] {
			// Point de départ
			if i.Nodes[y][x].Color.IsStartPoint() {
				start = i.Nodes[y][x]
			}

			// Point d'arrivée
			if i.Nodes[y][x].Color.IsEndPoint() {
				end = i.Nodes[y][x]
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

	defer i.findNeighbors()

	//Without go routines
	if (i.Width <= MIN_SIZE && i.Height <= MIN_SIZE) || (i.Width != i.Height) {
		i.Reader(0, i.Height)
	} else {
		//With go routines
		i.ReaderOptimized()
	}

	return i.Nodes

}

// Save image in ressources folder
func (i Image) Save(filename string) error {

	var img image.Image

	if (i.Width <= MIN_SIZE && i.Height <= MIN_SIZE) || (i.Width != i.Height) {
		img = i.Nodes.ToImage()
	} else {
		img = i.Nodes.ToImageOptimized()
	}

	out, err := os.Create("./ressources/" + filename)

	if err != nil {
		return err
	}

	defer out.Close()

	err = png.Encode(out, img)

	return err
}
