package Node

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
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

	//TODO Plus tard j'ajouterais des go routines, pour l'optimisation

	// Transforme l'image en tableau 2D de Node
	var nodes = make([][]*Node, height)

	for y := 0; y < height; y++ {
		var row = make([]*Node, width)
		for x := 0; x < width; x++ {

			// Extrait les valeurs RGB du pixel
			r, g, b, _ := img.At(x, y).RGBA()

			node := NewNode(x, y, NewColor(r, g, b))

			row[x] = node

		}
		nodes[y] = row
	}

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

func nodeToImage(nodes [][]*Node) image.Image {
	width, height := len(nodes), len(nodes[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := range nodes {
		for j := range nodes[i] {
			img.Set(j, i, color.RGBA{
				R: nodes[i][j].R,
				G: nodes[i][j].G,
				B: nodes[i][j].B,
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
