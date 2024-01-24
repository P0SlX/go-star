package main

import (
	"flag"
	"github.com/P0SLX/go-star/image"
	"github.com/P0SLX/go-star/utils"
	"log"
)

func main() {
	var imgPath string

	flag.StringVar(&imgPath, "img", "./ressources/pi.png", "Select path to image")
	flag.Parse()

	defer utils.Timer("main")()

	img, err := image.NewImage(imgPath)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	nodes := img.Read()

	start, end := img.FindStartAndEndNode(nodes)

	log.Printf("Start %#v, End %#v\n", start, end)

	/*	path := AStar.AStar(start, end)
		AStar.ColorPath(path)*/

	err = img.Save(nodes, "output.png")

	if err != nil {
		log.Fatalf("Error during image saving : %s\n", err.Error())
	}
}
