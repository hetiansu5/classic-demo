package main

import (
	"fmt"
)

//BM算法
func main() {
	offset := boyerMooreByBadCharRule([]rune("aaabaaabaaaabaaab"), []rune("aaaa"))
	fmt.Println(offset)

	offset2 := boyerMooreByGoodSuffixShift([]rune("aabcab"), []rune("abcab"))
	fmt.Println(offset2)

	offset3 := boyerMooreByBadCharRule([]rune("aaabaaabaaaabaaab"), []rune("aaaa"))
	fmt.Println(offset3)
}

//坏字符规则
//当发生不匹配的时候，我们把坏字符对应的模式串中的字符下标记作 si。
//如果坏字符在模式串中存在，我们把这个坏字符在模式串中的下标记作 xi。如果不存在，我们把 xi 记作 -1。那模式串往后移动的位数就等于 si-xi
func boyerMooreByBadCharRule(s []rune, sub []rune) int {
	n := len(s)
	m := len(sub)

	if m > n {
		return -1
	}

	offset := 0
	si := -1
	xi := -1
	for {
		if offset+m > n {
			return -1
		}

		si = -1
		for j := m - 1; j >= 0; j-- {
			//当发生不匹配的时候
			if s[offset+j] != sub[j] {
				si = j
				xi = getFinalIndexOfSub(s[offset+j], sub)
			}
		}
		if si == -1 {
			break
		} else if si-xi > 0 {
			offset += si - xi
		} else {
			offset += 1
		}
	}

	return offset
}

//可以使用散列表优化重复获取字符在子符串最后出现的位置
func getFinalIndexOfSub(r rune, sub []rune) int {
	l := len(sub)
	for i := l - 1; i >= 0; i-- {
		if sub[i] == r {
			return i
		}
	}
	return -1
}

//好后缀规则
func boyerMooreByGoodSuffixShift(s []rune, sub []rune) int {
	n := len(s)
	m := len(sub)

	if m > n {
		return -1
	}

	offset := 0
	suffix, prefix := generateGs(sub)
	for {
		if offset+m > n {
			return -1
		}

		j := m - 1
		for ; j >= 0; j-- {
			if s[offset+j] != sub[j] {
				break
			}
		}

		if j == - 1 {
			return offset
		} else if j < m-1 {
			offset += move(j, m, suffix, prefix)
		} else {
			offset += 1
		}
	}
}

//当存在好的后缀匹配时, 计算下一步的移动偏移量
//j表示坏字符串的位置
func move(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j //好后缀子串的长度
	if suffix[k] != -1 { //好后缀子串是否存在匹配的另子串
		return j + 1 - suffix[k] //好后缀子串的起始位置 - 另匹配子串的起始位置，即为移动偏移量
	}
	//如果不存在好后缀子串的另匹配子串，则好后缀子串的子串是否存在前缀匹配子串
	// j + 1 表示好后缀子串的起始位置， m - (j + 1) 表示好后缀子串的长度， m - (j + 2) 表示好后缀子串的子串的长度
	for r := j + 2; r < m; r++ {
		if prefix[m-r] == true {
			return r
		}
	}
	return m
}

//初始化模式串的子串匹配移动规则
func generateGs(sub []rune) (suffix []int, prefix []bool) {
	m := len(sub)
	suffix = make([]int, m)
	prefix = make([]bool, m)
	for i := 0; i < m-1; i++ {
		suffix[i] = -1
	}

	//每次从末位往前匹配，查找跟好后缀匹配的另一个子串，记为suffix，key为后缀子串的长度，value为另外一个子串的起始位置（-1表示不存在匹配）
	//prefix数组，key为后缀子串的长度， value表示是否存在匹配的前缀子串
	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for ; j >= 0 && sub[j] == sub[m-1-k]; {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}

	return
}

//结合坏字符规则和好后缀规则
func boyerMoore(s []rune, sub []rune) int {
	n := len(s)
	m := len(sub)

	if m > n {
		return -1
	}

	suffix, prefix := generateGs(sub)
	offset := 0
	si := -1
	xi := -1
	for {
		if offset+m > n {
			return -1
		}

		si = -1
		j := m - 1
		for ; j >= 0; j-- {
			//当发生不匹配的时候
			if s[offset+j] != sub[j] {
				si = j
				xi = getFinalIndexOfSub(s[offset+j], sub)
			}
		}


		if si == -1 {
			break
		}

		x := si - xi
		y := 0
		if j < m - 1 {
			y = move(j, m, suffix, prefix)
		}

		if x > y {
			y = x
		}
		offset += y
	}

	return offset
}
