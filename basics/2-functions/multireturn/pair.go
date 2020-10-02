package multireturn

func ReturnPair(x, y int) (int, int) {
	return x + y, x * y
}

func ReturnPairFancy(x, y int) (sum int, multi int) {
	sum = x + y
	multi = x * y

	return
}
