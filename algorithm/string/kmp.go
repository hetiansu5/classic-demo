package main

import "fmt"

func main() {
	offset := KMP([]rune("aaabaaabaaaabaaab"), []rune("ababc"))
	fmt.Println(genNext([]rune("ababc")))
}

func KMP(str []rune, sub []rune) int {
	next := genNext(sub)
	fmt.Println(next)
	n := len(str)
	m := len(sub)
	var k int
	for i := 0; i < n; {
		if str[i+k] == sub[k] {
			if k == m-1 {
				return i
			}
			k++
		} else {
			if k == 0 {
				i++
			} else {
				i += next[k-1]
				k = 0
			}
		}
	}
	return -1
}

func generateNext(sub []rune) []int {
	m := len(sub)
	next := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		next[i] = -1
		p := i
		for ; p >= 1; p-- {
			j := p - 1
			k := 0
			for ; j >= 0 && sub[j] == sub[p-k]; {
				j--
				k++
			}
			if j == -1 {
				next[i] = m - p
				break
			}
		}
	}
	return next
}

//next 模式串前缀串长度（假定主串与模式串匹配的长度） value为模式串前缀串的最大前后缀匹配长度
func genNext(sub []rune) []int {
	m := len(sub)
	next := make([]int, m)
	next[0] = -1
	i := 0
	j := -1
	for ; i < m - 1; {
		if j == -1 || sub[i] == sub[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j] //次最大匹配长度
		}
	}
	return next
}
