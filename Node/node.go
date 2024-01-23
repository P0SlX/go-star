package Node

import (
	"image"
	"io"
	"sync"
)

type Node struct {
	X int
	Y int

	*Color
}

func NewNode(x, y int, color *Color) *Node {
	return &Node{
		x,
		y,
		color,
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

	return nodes, nil
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
			if nodes[i][j].isEndPoint() {
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
