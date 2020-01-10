package standard

import "fmt"

//将slice或者map数据作为参数传递时，需要注意先copy再操作，这是因为slice和map的底层数据结构设计的问题。
func  SetTrips(trips []int) {
	tripsCopy := make([]int, len(trips))
	copy(tripsCopy, trips)
	fmt.Println(trips)
}
