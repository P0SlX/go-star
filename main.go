package main

import (
	"flag"
	"fmt"
	"github.com/P0SLX/go-star/image"
	"github.com/P0SLX/go-star/node"
	"github.com/P0SLX/go-star/utils"
	"log"
)

func main() {

	var imgPath string

	flag.StringVar(&imgPath, "img", "./ressources/first_level.png", "Select path to image")
	flag.Parse()

	defer utils.Timer("main")()

	data, err := image.NewImage(imgPath)

	nodes, err := node.GetNodes(file)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	start, end := node.GetStartAndEnd(nodes)

	fmt.Printf("Start %#v, End %#v\n", start, end)

	err = node.SaveToFile(nodes, "./output.png")

	if err != nil {
		log.Fatalf("Error during image saving : %s\n", err.Error())
	}
}
