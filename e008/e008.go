package e008

import (
	"fmt"
	"strconv"
	"strings"
)

// s - String repr. of number
// c - Count, size of sequence
func NaiveProdSeq(s string, c int) int {
	r := []rune(s)
	max := 0

	if c <= 0 {
		c = 4
	}

	for i := 0; i < len(s)-c; i++ {
		// Product
		p := 1

		// Skip sequences with a zero
		if strings.Contains(string(r[i:i+c]), "0") {
			continue
		}

		for j := 0; j < c; j++ {
			// Current digit
			d, err := strconv.Atoi(string(r[i+j]))

			if err != nil {
				fmt.Printf("error: converting %d to int failed\n", d)
				return 0
			} else {
				p *= d
			}
		}

		if p > max {
			max = p
		}
	}

	return max
}
