package factors003_test

import (
	"reflect"
	"testing"

	f "github.com/nevilleomangi/euler/factors003"
)

func TestPrimeFactors(t *testing.T) {
	got := f.PrimeFactorsTrial(13195)
	want := []int{5, 7, 13, 29}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
