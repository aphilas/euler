package e005_test

import (
	"testing"

	"github.com/nevilleomangi/euler/e005"
)

func TestRangeLCMNaive(t *testing.T) {
	got := e005.RangeLCMNaive(10)
	want := 2520

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}

func TestRangeLCMEuclid(t *testing.T) {
	got := e005.RangeLCMEuclid(10)
	want := 2520

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}

func BenchmarkRangeLCMNaive(b *testing.B) {
	e005.RangeLCMNaive(20)
}

func BenchmarkRangeLCMEuclid(b *testing.B) {
	e005.RangeLCMEuclid(20)
}
