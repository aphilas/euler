package e007

import (
	"math"

	"github.com/nevilleomangi/euler/e003"
)

// Rosser, 1962?
// Works for n >= 6
func PrimeNLim(n float64) float64 {
	return n*math.Log(n) + n*math.Log(math.Log(n))
}

func PrimeN(n int) int {
	// 5th prime is 11
	limit := 11

	if n >= 6 {
		limit = int(PrimeNLim(float64(n)))
	}

	primes := e003.SieveOfEratosthenes(limit)
	return primes[n-1]
}
