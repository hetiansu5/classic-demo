package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "1"
	m[30] = "30"
	//map遍历的语法糖
	// Lower a for range over a map.
	// The loop we generate:
	//   var hiter map_iteration_struct
	//   for mapiterinit(type, range, &hiter); hiter.key != nil; mapiternext(&hiter) {
	//           index_temp = *hiter.key
	//           value_temp = *hiter.val
	//           index = index_temp
	//           value = value_temp
	//           original body
	//   }
	for k := range m {
		fmt.Println(k)
		if k > 10 {
			m[2] = "2"
		}
	}
}


//slice遍历的语法糖
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }

//channel
// Lower a for range over a channel.
// The loop we generate:
//   for {
//           index_temp, ok_temp = <-range
//           if !ok_temp {
//                   break
//           }
//           index = index_temp
//           original body
//   }

//array
// Lower a for range over an array.
// The loop we generate:
//   len_temp := len(range)
//   range_temp := range
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = range_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }


//string
// Lower a for range over a string.
// The loop we generate:
//   len_temp := len(range)
//   var next_index_temp int
//   for index_temp = 0; index_temp < len_temp; index_temp = next_index_temp {
//           value_temp = rune(range[index_temp])
//           if value_temp < utf8.RuneSelf {
//                   next_index_temp = index_temp + 1
//           } else {
//                   value_temp, next_index_temp = decoderune(range, index_temp)
//           }
//           index = index_temp
//           value = value_temp
//           // original body
//   }