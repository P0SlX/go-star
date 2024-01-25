package utils

import (
	"log"
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
		log.Printf("%s() %v\n", name, time.Since(start))
	}
}

func FindBestChunck(height int) (chunck int, rest int, resolution int) {
	var maxCPU = runtime.NumCPU() * 10

	var resolutions = []int{4, 8, 16, 32, 64, 128, 256, 512, 1_024, 2_048, 4_096, 8_192, 16_384, 32_768}

	for _, r := range resolutions {
		chunck = height / r

		rest = height % r

		if chunck < maxCPU {
			return chunck, rest, r
		}
	}

	return chunck, rest, resolution

}
