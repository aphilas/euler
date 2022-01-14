package e006

func SumSquares(stop int) int {
	return (stop * (stop + 1) * (2*stop + 1)) / 6
}

func SquareSum(stop int) int {
	sum := int((1 + float32(stop)) / 2 * float32(stop))
	return sum * sum
}
