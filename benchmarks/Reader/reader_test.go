package Reader

import (
	"github.com/P0SLX/go-star/benchmarks"
	_ "log"
	"testing"
)

func BenchmarkSmallImage(b *testing.B) {
	img, height := benchmarks.InitBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		img.Reader(0, height)
	}
}

func BenchmarkMediumImage(b *testing.B) {
	img, height := benchmarks.InitBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		img.Reader(0, height)
	}
}

func BenchmarkLargeImage(b *testing.B) {
	img, height := benchmarks.InitBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		img.Reader(0, height)
	}
}

func BenchmarkSmallImageOptimized(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/first_level.png")

	for n := 0; n < b.N; n++ {
		img.ReaderOptimized()
	}
}

func BenchmarkMediumImageOptimized(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/pi.png")

	for n := 0; n < b.N; n++ {
		img.ReaderOptimized()
	}
}

func BenchmarkLargeImageOptimized(b *testing.B) {
	img, _ := benchmarks.InitBenchmark("../../ressources/large.png")

	for n := 0; n < b.N; n++ {
		img.ReaderOptimized()
	}
}
