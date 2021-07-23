package main

import "fmt"

func main() {

}

func Emav2() {
	// on cherche à calculer l'EMA en base 10 au jour 50
	// EMA(T) = price(T) * multiplicateur + ema(T-1)*(1-multiplicateur)
	// multiplicateur = 2/(days+1)

	// sample := []float32{12, 13, 15, 18, 17, 18, 20, 21, 24, 24, 31, 21, 25, 20, 21, 50}
	// base := float32(10)
	// multi := 2 / (base + 1)
	// EMA := sample[len(sample)-1] * multi +

}

func Ema() {
	N := 10
	sample := []float32{12, 13, 15, 18, 17, 18, 20, 21, 24, 24, 31, 21, 25, 20, 21, 50}
	var constanteDePonderation float32

	movingAverageN := moyenne(sample[len(sample)-(N):])
	constanteDePonderation = float32(2) / float32(N+1)
	fmt.Println(movingAverageN)
	fmt.Println(constanteDePonderation)

	EMANP1 := ((sample[len(sample)-1] - movingAverageN) * constanteDePonderation) + movingAverageN
	fmt.Println(EMANP1)
}

func moyenne(liste []float32) float32 {
	fmt.Println(liste)
	var sum float32

	for _, v := range liste {
		sum += v
	}
	return sum / float32(len(liste))
}
