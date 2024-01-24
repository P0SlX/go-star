package main

import (
	"flag"
	"github.com/P0SLX/go-star/AStar"
	"github.com/P0SLX/go-star/image"
	"log"
)

func main() {
	var imgPath string

	flag.StringVar(&imgPath, "img", "./ressources/first_level.png", "Select path to image")
	flag.Parse()

	img, err := image.NewImage(imgPath)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	nodes := img.Read()

	start, end := img.FindStartAndEndNode(nodes)

	path := AStar.AStar(start, end)
	AStar.ColorPath(path)

	err = img.Save(nodes, "output.png")

	if err != nil {
		log.Fatalf("Error during image saving : %s\n", err.Error())
	}
}
