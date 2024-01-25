package Save

import (
	"github.com/P0SLX/go-star/benchmarks"
	_ "log"
	"testing"
)

func BenchmarkSaveSmallImage(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		img.Nodes.ToImage()
	}
}

func BenchmarkSaveMediumImage(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		img.Nodes.ToImage()
	}
}

func BenchmarkSaveLargeImage(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/xxl.png")

	for n := 0; n < b.N; n++ {
		img.Nodes.ToImage()
	}
}

func BenchmarkSaveSmallImageOptimize(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		img.Nodes.ToImageOptimized()
	}
}

func BenchmarkSaveMediumImageOptimize(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		img.Nodes.ToImageOptimized()
	}
}

func BenchmarkSaveLargeImageOptimize(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/xxl.png")

	for n := 0; n < b.N; n++ {
		img.Nodes.ToImageOptimized()
	}
}
