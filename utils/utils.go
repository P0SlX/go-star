package utils

import (
	"fmt"
	"github.com/P0SLX/go-star/node"
	"image"
	"image/color"
	"runtime"
	"time"
)

// Timer permet de mesurer le temps d'exécution d'une fonction jusqu'à son retour
//
// Usage : defer timer("nomDeLaFonctionAMesurer")()
//
// Source : https://stackoverflow.com/a/45766707
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s() %v\n", name, time.Since(start))
	}
}

func FindBestChunck(height int) (int, int, int) {
	var maxCPU = runtime.NumCPU() * 10

	var resolutions = []int{4, 8, 16, 32, 64, 128, 256, 512, 1_024, 2_048, 4_096, 8_192, 16_384, 32_768}

	var chunck, rest int

	for _, resolution := range resolutions {
		chunck = height / resolution

		rest = height % resolution

		if chunck < maxCPU {
			return chunck, rest, resolution
		}
	}

	return 0, 0, 0

}

func NodeToImageOptimize(nodes [][]*node.Node) image.Image {
	width, height := len(nodes), len(nodes[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	chunk, rest, factor := FindBestChunck(height)

	maxWorkers := runtime.NumCPU()

	sem := make(chan struct{}, maxWorkers)

	for j := 0; j < chunk; j++ {
		sem <- struct{}{}
		go func(width, start, end int, nodes [][]*node.Node) {
			defer func() { <-sem }()
			for y := start; y < end; y++ {
				for x := 0; x < width; x++ {
					img.Set(y, x, color.RGBA{
						R: nodes[x][y].Color.R,
						G: nodes[x][y].Color.G,
						B: nodes[x][y].Color.B,
						A: 255,
					})
				}
			}

		}(width, j*factor, (j+1)*factor, nodes)
	}

	for j := 0; j < maxWorkers; j++ {
		sem <- struct{}{}
	}

	if rest > 0 {
		for i := chunk * factor; i < chunk*factor+rest; i++ {
			for k := range nodes[i] {
				img.Set(k, i, color.RGBA{
					R: nodes[i][k].Color.R,
					G: nodes[i][k].Color.G,
					B: nodes[i][k].Color.B,
					A: 255,
				})
			}
		}
	}

	return img
}

func NodeToImage(nodes [][]*node.Node) image.Image {
	width, height := len(nodes), len(nodes[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := range nodes {
		for j := range nodes[i] {
			img.Set(j, i, color.RGBA{
				R: nodes[i][j].Color.R,
				G: nodes[i][j].Color.G,
				B: nodes[i][j].Color.B,
				A: 255,
			})
		}
	}
	return img
}
