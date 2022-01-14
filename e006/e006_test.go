package e006_test

import (
	"testing"

	"github.com/nevilleomangi/euler/e006"
)

func TestSumSquares(t *testing.T) {
	got := e006.SumSquares(10)
	want := 385

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}

func TestSquareSum(t *testing.T) {
	got := e006.SquareSum(10)
	want := 3025

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
