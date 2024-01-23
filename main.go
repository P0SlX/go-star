package main

import (
	"fmt"
	"github.com/P0SLX/go-star/Node"
	"image"
	"image/png"
	"os"
)

func main() {
	defer Timer("main")()
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, fileErr := os.Open("./first_level.png")

	if fileErr != nil {
		fmt.Println("Image introuvable...")
		os.Exit(1)
	}

	//Pas obligé, juste le defer de file.Close devrait suffire
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Impossible de fermer le fichier")
			os.Exit(1)
		}
	}(file)

	nodes, nodesErr := Node.GetNodes(file)

	if nodesErr != nil {
		fmt.Println("Impossible de décoder l'image")
		os.Exit(1)
	}

	start, end := Node.GetStartAndEnd(nodes)

	fmt.Println(start, end)

	Node.SaveToFile(nodes, "./output.png")
}
