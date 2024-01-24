package Node

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"sync"
)

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

func loopOverImages(img image.Image, width, start, end int, nodes *[][]*Node) {

	for y := start; y < end; y++ {
		var row = make([]*Node, width)
		for x := 0; x < width; x++ {

			// Extrait les valeurs RGB du pixel
			r, g, b, _ := img.At(x, y).RGBA()

			node := NewNode(x, y, NewColor(r, g, b))

			row[x] = node

		}
		(*nodes)[y] = row
	}

}

// GetNodes Extrait les pixels d'une image en un tableau 2D de Node
//
// La fonction décode l'image, obtient les limites,
// initialise un tableau 2D de Node, et boucle sur chaque pixel.
func GetNodes(file io.Reader) ([][]*Node, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	// Limites de l'image
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var nodes = make([][]*Node, height)

	//Without go routines
	if width <= 16 && height <= 16 {
		loopOverImages(img, width, 0, height, &nodes)
		applyNeighbors(nodes)

		return nodes, nil
	}

	//With go routines
	//TODO : Trouver un moyen de ne pas hardcoder le nombre de go routines, plutôt proportionnel à la taille de l'image
	chunck := height / 10
	wg := sync.WaitGroup{}
	wg.Add(chunck)

	for i := 0; i < chunck; i++ {
		go func(img image.Image, width, start, end int, nodes *[][]*Node) {
			defer wg.Done()
			loopOverImages(img, width, start, end, nodes)
		}(img, width, i*10, (i+1)*10-1, &nodes)
	}

	wg.Wait()

	applyNeighbors(nodes)
	return nodes, nil
}

func applyNeighbors(nodes [][]*Node) {
	for i := range nodes {
		for j := range nodes[i] {
			// Haut
			if i > 0 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i-1][j])
			}

			// Bas
			if i < len(nodes)-1 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i+1][j])
			}

			// Gauche
			if j > 0 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i][j-1])
			}

			// Droite
			if j < len(nodes[i])-1 {
				nodes[i][j].Neighbors = append(nodes[i][j].Neighbors, nodes[i][j+1])
			}
		}
	}
}

// GetStartAndEnd Détecte le point de départ et d'arrivée
//
// Le point de départ est un pixel vert (0, 255, 0)
// Le point d'arrivée est un pixel rouge (255, 0, 0)
// On boucle sur chaque node, et on renvoie le pointeur
// du premier pixel vert et du premier pixel rouge
func GetStartAndEnd(nodes [][]*Node) (*Node, *Node) {
	var start *Node
	var end *Node

	for i := range nodes {
		for j := range nodes[i] {
			// Point de départ
			if nodes[i][j].Color.isStartPoint() {
				start = nodes[i][j]
			}

			// Point d'arrivée
			if nodes[i][j].Color.isEndPoint() {
				end = nodes[i][j]
			}

			// Tout trouvé ? Pas besoin de continuer
			if start != nil && end != nil {
				return start, end
			}
		}
	}

	return nil, nil
}

// ColorPath Colorie le chemin trouvé en violet
func ColorPath(nodes []*Node) {
	for _, node := range nodes {
		node.Color.R = 255
		node.Color.G = 0
		node.Color.B = 255
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

func SaveToFile(nodes [][]*Node, filename string) {
	img := nodeToImage(nodes)

	out, saveErr := os.Create(filename)
	if saveErr != nil {
		fmt.Println("Impossible de créer le fichier de sortie")
		os.Exit(1)
	}
	defer func(out *os.File) {
		closeErr := out.Close()
		if closeErr != nil {

		}
	}(out)

	encodeErr := png.Encode(out, img)
	if encodeErr != nil {
		return
	}
}
