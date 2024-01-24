package Reader

import (
	"github.com/P0SLX/go-star/image"
	"github.com/P0SLX/go-star/node"
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

func BenchmarkSmallImage(b *testing.B) {
	img, width, height, nodes := initBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		img.Reader(width, 0, height, &nodes)
	}
}

func BenchmarkMediumImage(b *testing.B) {
	img, width, height, nodes := initBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		img.Reader(width, 0, height, &nodes)
	}
}

func BenchmarkLargeImage(b *testing.B) {
	img, width, height, nodes := initBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		img.Reader(width, 0, height, &nodes)
	}
}

func BenchmarkSmallImageOptimized(b *testing.B) {
	img, width, height, nodes := initBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		img.ReaderOptimized(width, height, &nodes)
	}
}

func BenchmarkMediumImageOptimized(b *testing.B) {
	img, width, height, nodes := initBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		img.ReaderOptimized(width, height, &nodes)
	}
}

func BenchmarkLargeImageOptimized(b *testing.B) {
	img, width, height, nodes := initBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		img.ReaderOptimized(width, height, &nodes)
	}
}
