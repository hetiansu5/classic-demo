package main

func main()  {
	getArr(45)
}

func getArr(a int) int {
	var arr = [...]int{1 , 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	return arr[a]
}
