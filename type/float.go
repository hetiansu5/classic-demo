package main

import (
	"strconv"
	"fmt"
	"encoding/json"
)

func main()  {
	data := "{\"id\": 1}11";
	v := json.Valid([]byte(data))
	fmt.Println(v)

	var f = float64(2.10100)
	fmt.Println(f)
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 64))
}
