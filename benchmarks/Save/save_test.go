package Save

import (
	"github.com/P0SLX/go-star/benchmarks"
	"github.com/P0SLX/go-star/utils"
	_ "log"
	"testing"
)

func BenchmarkSaveSmallImage(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImage(img.Nodes)
	}
}

func BenchmarkSaveMediumImage(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImage(img.Nodes)
	}
}

func BenchmarkSaveLargeImage(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImage(img.Nodes)
	}
}

func BenchmarkSaveSmallImageOptimize(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImageOptimize(img.Nodes)
	}
}

func BenchmarkSaveMediumImageOptimize(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImageOptimize(img.Nodes)
	}
}

func BenchmarkSaveLargeImageOptimize(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		utils.NodeToImageOptimize(img.Nodes)
	}
}
