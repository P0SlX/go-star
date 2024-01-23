package main

import (
	"fmt"
	"time"
)

// Fonction de mesure de temps
// Permet de mesurer le temps d'exécution d'une fonction jusqu'à son retour
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
