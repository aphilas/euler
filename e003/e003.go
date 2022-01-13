package e003

import "math"

func SieveOfEratosthenes(n int) []int {
	// Start from 0
	sieve := make([]bool, n+1)

	// Mark 0 and 1 as composites
	sieve[0], sieve[1] = true, true

	for p := 2; p*p <= n; p++ {
		if !sieve[p] {
			for i := p * 2; i <= n; i += p {
				sieve[i] = true
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if !sieve[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func PrimeFactorsTrial(n int) []int {
	factors := []int{}
	primes := SieveOfEratosthenes(int(math.Ceil(math.Sqrt(float64(n)))))

	for _, p := range primes {
		for n%p == 0 {
			factors = append(factors, p)
			n /= p
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
