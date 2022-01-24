package e010_test

import (
	"testing"

	"github.com/nevilleomangi/euler/e010"
)

func TestSumPrime(t *testing.T) {
	const limit int = 10
	want := 17
	got := e010.SumPrimes(limit)

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
