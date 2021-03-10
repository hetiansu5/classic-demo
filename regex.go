package main

import (
	"fmt"
	"strings"
)

const (
	FlagMiss  = 0
	FlagMatch = 1
)

func main() {
	var str = "select * from user where mobile = ${mobile} and phone in (${phone})"
	l := len(str)

	var sqlBuilder strings.Builder
	var keyArr []string
	var keyStart, sqlStart int
	var flag uint8
	for i := 0; i < l; i++ {
		if str[i] == '$' && i+1 < l && str[i+1] == '{' {
			flag = FlagMatch
			sqlBuilder.WriteString(str[sqlStart:i])
			keyStart = i + 2
			i++
			continue
		}

		if flag == FlagMatch {
			if str[i] == '}' {
				sqlBuilder.WriteString("?")
				keyArr = append(keyArr, str[keyStart:i])
				flag = FlagMiss
				sqlStart = i + 1
			}
		}
	}

	if sqlStart < l {
		sqlBuilder.WriteString(str[sqlStart:l])
	}

	sql := sqlBuilder.String()

	fmt.Println(sql)
	fmt.Println(keyArr)
}

