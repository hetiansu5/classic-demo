package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "吉1a人ü§䄼힞Ⱥ😃"
	fmt.Printf("The string: %q\n", str)

	// 当 string 转换为 []rune 时，字符串会被拆分成一个个 Unicode 字符
	fmt.Printf(" => runs(char): %q\n", []rune(str))

	// rune 在底层由 UTF-8 编码值表达
	// 英文字符的 UTF-8 编码由一个字节表达
	// 中文字符的 UTF-8 编码由三个字节表达
	fmt.Printf(" => runs(hex): %x\n", []rune(str))

	// 进一步，把每个字符的 UTF-8 编码值拆成相应的字节序列
	fmt.Printf(" => bytes(hex): [% x]\n", []byte(str))

	// i 为索引，c 为 UTF-8 编码值代表的 Unicode 字符，类型为 rune
	// 相邻的 Unicode 字符的索引值不一定连续，取决于前一个 Unicode 字符是否为单字节字符
	for i, c := range str {
		bytes := []byte(string(c))
		fmt.Printf("%d: %q %d %b [% x] %b \n", i, c, c, c, bytes, bytes)
		if !checkCodePointToUTF8ByteRule(string(c)) {
			fmt.Println("errr:", string(c))
		}
	}
}

//码点与utf-8字节二进制的转换规则
func checkCodePointToUTF8ByteRule(s string) bool {
	r := []rune(s)[0]
	d := int64(r);
	binary := strconv.FormatInt(d, 2)
	l := len(binary)

	var res int64
	if d <= 0x007F {
		res = d
	} else if d <= 0x07FF {
		res, _ = strconv.ParseInt("110" + fmt.Sprintf("%05s", binary[0:l-6]) + "10" + binary[l - 6:], 2, 64)
	} else if d <= 0xFFFF {
		res, _ = strconv.ParseInt("1110" + fmt.Sprintf("%04s", binary[0:l-12]) + "10" + binary[l-12:l-6] +
			"10" + binary[l - 6:], 2, 64)
	} else {
		var (
			pre string
			mid string
		)
		if l >= 18 {
			pre = fmt.Sprintf("%03s", binary[0:l-18])
			mid = binary[l-18:l-12]
		} else {
			pre = "000"
			mid = fmt.Sprintf("%06s", binary[0:l-12])
		}
		res, _ = strconv.ParseInt("11110" + pre + "10" + mid + "10" + binary[l-12:l-6] + "10" + binary[l - 6:], 2, 64)
	}
	//fmt.Println(d)
	//fmt.Println(binary)
	//fmt.Println(strconv.FormatInt(res, 2))
	//fmt.Println(fmt.Sprintf("%x", []byte(s)), strconv.FormatInt(res, 16))
	return fmt.Sprintf("%x", []byte(s)) == strconv.FormatInt(res, 16)
}
