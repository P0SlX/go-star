package benchmarks

import (
	"github.com/P0SLX/go-star/image"
	"log"
)

func InitBenchmark(path string) (*image.Image, int) {
	img, err := image.NewImage(path)

	if err != nil {
		log.Fatalf("Error during image decoding : %s\n", err.Error())
	}

	return img, img.Height
}
