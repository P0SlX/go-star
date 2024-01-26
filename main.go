package main

import (
	"flag"
	"github.com/P0SLX/go-star/astar"
	"github.com/P0SLX/go-star/image"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	var imgPath string

	flag.StringVar(&imgPath, "img", "./ressources/xxl.png", "Select path to image")
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	time.Sleep(15 * time.Second)

	img, err := image.NewImage(imgPath)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	img.Read()

	start, end := img.FindStartAndEndNode()
	path := astar.AStar(start, end)
	astar.ColorPath(path)

	err = img.Save("output.png")

	if err != nil {
		log.Fatalf("Error during image saving : %s\n", err.Error())
	}

	time.Sleep(1 * time.Minute)

}
