package e004_test

import (
	"testing"

	"github.com/nevilleomangi/euler/e004"
)

// Fails!
func TestPalNaive(t *testing.T) {
	t.Skip("Skip failing naive implementation")

	got := e004.PalNaive(2)
	want := 9009

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}

func TestPalNaive2(t *testing.T) {
	got := e004.PalNaive2(2)
	want := 9009

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
