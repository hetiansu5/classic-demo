package main

func f(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return f(i - 1)
}

func main() {
	println(f(9999))
}
