package Save

import (
	"github.com/P0SLX/go-star/image"
	"github.com/P0SLX/go-star/node"
	"github.com/P0SLX/go-star/utils"
	"log"
	_ "log"
	"testing"
)

func initBenchmark(path string) (*image.Image, int, int, [][]*node.Node) {
	img, err := image.NewImage(path)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	bounds := img.Image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	return img, width, height, make([][]*node.Node, height)
}

func BenchmarkSaveSmallImage(b *testing.B) {
	_, _, _, nodes := initBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImage(nodes)
	}
}

func BenchmarkSaveMediumImage(b *testing.B) {
	_, _, _, nodes := initBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImage(nodes)
	}
}

func BenchmarkSaveLargeImage(b *testing.B) {
	_, _, _, nodes := initBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImage(nodes)
	}
}

func BenchmarkSaveSmallImageOptimize(b *testing.B) {
	_, _, _, nodes := initBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImageOptimize(nodes)
	}
}

func BenchmarkSaveMediumImageOptimize(b *testing.B) {
	_, _, _, nodes := initBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImageOptimize(nodes)
	}
}

func BenchmarkSaveLargeImageOptimize(b *testing.B) {
	_, _, _, nodes := initBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImageOptimize(nodes)
	}
}
