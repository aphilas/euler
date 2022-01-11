package main

import (
	"fmt"

	"github.com/nevilleomangi/euler/factors003"
)

func main() {
	n := 600851475143
	f := factors003.PrimeFactorsTrial(n)
	fmt.Println(f[len(f)-1])
}
