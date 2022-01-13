package e005

func RangeLCMNaive(stop int) int {
	for n := stop; ; n++ {
		flag := true

		for i := 2; i <= stop; i++ {
			if n%i != 0 {
				flag = false
				break
			}
		}

		if flag {
			return n
		}
	}
}

func EuclidGCD(a, b int) int {
	var r int

	if a < b {
		a, b = b, a
	}

	for {
		r = a % b
		a, b = b, r

		if r == 0 {
			return a
		}
	}
}

func EuclidLCM(a, b int) int {
	return a / EuclidGCD(a, b) * b
}

func RangeLCMEuclid(stop int) int {
	acc := 1

	for i := 2; i <= stop; i++ {
		acc = EuclidLCM(acc, i)
	}

	return acc
}
