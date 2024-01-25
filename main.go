package main

import (
	"flag"
	"github.com/P0SLX/go-star/image"
	"github.com/P0SLX/go-star/utils"
	"log"
)

func main() {
	var imgPath string

	flag.StringVar(&imgPath, "img", "./ressources/large.png", "Select path to image")
	flag.Parse()

	defer utils.Timer("main")()

	img, err := image.NewImage(imgPath)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	_ = img.Read()

	/*
		start, end := img.FindStartAndEndNode()
		path := astar.astar(start, end)
		astar.ColorPath(path)
	*/

	err = img.Save("output.png")

	if err != nil {
		log.Fatalf("Error during image saving : %s\n", err.Error())
	}
}
