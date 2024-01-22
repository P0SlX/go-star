package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"time"
)

type Node struct {
	X int
	Y int

	R int
	G int
	B int
}

// Fonction de mesure de temps
// Permet de mesurer le temps d'exécution d'une fonction jusqu'à son retour
//
// Usage : defer timer("nomDeLaFonctionAMesurer")()
//
// Source : https://stackoverflow.com/a/45766707
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s() %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, fileErr := os.Open("./first_level.png")

	if fileErr != nil {
		fmt.Println("Image introuvable...")
		os.Exit(1)
	}

	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Impossible de fermer le fichier")
			os.Exit(1)
		}
	}(file)

	nodes, nodesErr := getNodes(file)

	if nodesErr != nil {
		fmt.Println("Impossible de décoder l'image")
		os.Exit(1)
	}

	start, end := getStartAndEnd(nodes)

	fmt.Println(start, end)
}

// Extrait les pixels d'une image en un tableau 2D de Node
//
// La fonction décode l'image, obtient les limites,
// initialise un tableau 2D de Node, et boucle sur chaque pixel.
func getNodes(file io.Reader) ([][]Node, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	// Limites de l'image
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Transforme l'image en tableau 2D de Node
	var nodes [][]Node
	for y := 0; y < height; y++ {
		var row []Node
		for x := 0; x < width; x++ {
			// Extrait les valeurs RGB du pixel
			r, g, b, _ := img.At(x, y).RGBA()

			// Divise par 257 pour obtenir des valeurs entre 0 et 255
			node := Node{
				x,
				y,
				int(r / 257),
				int(g / 257),
				int(b / 257),
			}
			row = append(row, node)

		}
		nodes = append(nodes, row)
	}

	return nodes, nil
}

// Détecte le point de départ et d'arrivée
//
// Le point de départ est un pixel vert (0, 255, 0)
// Le point d'arrivée est un pixel rouge (255, 0, 0)
// On boucle sur chaque node, et on renvoie le pointeur
// du premier pixel vert et du premier pixel rouge
func getStartAndEnd(nodes [][]Node) (*Node, *Node) {
	var start *Node
	var end *Node

	for i := range nodes {
		for j := range nodes[i] {
			// Point de départ
			if nodes[i][j].R == 0 && nodes[i][j].G == 255 && nodes[i][j].B == 0 {
				start = &nodes[i][j]
			}

			// Point d'arrivée
			if nodes[i][j].R == 255 && nodes[i][j].G == 0 && nodes[i][j].B == 0 {
				end = &nodes[i][j]
			}

			// Tout trouvé ? Pas besoin de continuer
			if start != nil && end != nil {
				return start, end
			}
		}
	}

	return nil, nil
}
