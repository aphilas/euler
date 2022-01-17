package e004

// Find the largest palindromic number that is a product of 3-digit integers

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Reverses a string
func RevStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Reverses an int
func RevInt(value int) (int, error) {
	s := strconv.Itoa(value)
	r := RevStr(s)
	v, err := strconv.Atoi(r)

	if err != nil {
		return 0, errors.New("cannot reverse int")
	}

	return v, nil
}

// Fails!
func PalNaive(size int) int {
	p := int(math.Pow10(size) - 1)

	m, n := p, p
	f := true // Flag

	for {
		p := m * n
		r, err := RevInt(p)

		if err != nil {
			fmt.Println(err)
			return 0
		}

		if p == r {
			return p
		}

		if f {
			m -= 1
		} else {
			n -= 1
		}

		f = !f
	}
}

func PalNaive2(size int) int {
	a, b := int(math.Pow10(size-1)), int(math.Pow10(size)-1)
	max := 0

	for m := b; m > a; m -= 1 {
		for n := m; n > a; n-- {
			p := m * n
			r, err := RevInt(p)

			if err != nil {
				fmt.Println(err)
				return 0
			}

			if p == r && p > max {
				max = p
			}
		}
	}

	return max
}
