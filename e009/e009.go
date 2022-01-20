package e009

// Pythagorean triple - a^2 + b^2 = c^2
// If a+b+c = 1000, find a*b*c

func EuclideanTriples(lim int) int {
	for m := 2; ; m++ {
		for n := m%2 + 1; n < m; n++ {
			for k := 1; ; k++ {
				a, b, c := k*(m*m-n*n), k*2*m*n, k*(m*m+n*n)
				sum := a + b + c

				if sum > 1000 {
					if k == 1 {
						return 0
					}

					break
				} else if sum == 1000 {
					return a * b * c
				}
			}
		}
	}
}
