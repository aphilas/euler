package e007_test

import (
	"reflect"
	"testing"

	"github.com/nevilleomangi/euler/e007"
)

func TestPrimeN(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17, 19}
	c := len(want)

	got := make([]int, c)
	for i := 1; i <= c; i++ {
		got[i-1] = e007.PrimeN(i)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
