// Find the sum of all primes below two million

package e010

import (
	"github.com/nevilleomangi/euler/e003"
)

func Sum(v []int) (sum int) {
	for _, n := range v {
		sum += n
	}
	return
}

func SumPrimes(limit int) int {
	primes := e003.SieveOfEratosthenes(limit)
	return Sum(primes)
}
