package main

func sumArray(a byte, b int, c int) (int, int, int, int) {
	var d int
	var f int
	d = b * 2
	f, _ = addOne(int(a), b, c)
	return int(a) + b + c + d + f, 0, 0, 0
}

func addOne(m, n, k int) (int, int) {
	return addTwo(m, m), n + k
}

func addTwo(o, p int) int{
	return o + p
}
