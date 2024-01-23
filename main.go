package main

import (
	"flag"
	"fmt"
	"github.com/P0SLX/go-star/Node"
	"image"
	"image/png"
	"os"
)

func main() {

	var imgPath string

	flag.StringVar(&imgPath, "img", "./ressources/first_level.png", "Select path to image")
	flag.Parse()

	//defer Timer("main")()
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open(imgPath)

	if err != nil {
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

	nodes, err := Node.GetNodes(file)

	if err != nil {
		fmt.Println("Impossible de décoder l'image")
		os.Exit(1)
	}

	start, end := Node.GetStartAndEnd(nodes)

	fmt.Printf("Start %#v, End %#v\n", start, end)
}
