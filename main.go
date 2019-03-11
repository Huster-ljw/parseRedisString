package main

import (
	"fmt"
	"strings"
)

/************************************************************************/
/*    第一个字节是“+”代表简单字符串类型。     "+OK\r\n"
/*    第一个字节是“-”代表错误类型。           "-Error message\r\n"
/*    第一个字节是“:”代表整型。               ":1000\r\n"
/*    第一个字节是“$”代表块字符串。           "$6\r\nfoobar\r\n"
/*    第一个自己是“*”代表数组。               "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
/************************************************************************/
func parseString(str string) (strs []string) {
	if str == "" {
		panic("input error")
	}

	if str[0] == '+' || str[0] == '-' || str[0] == ':' {
		spls := strings.Split(str, "\r\n")
		strs = append(strs, spls[0][1:])
	} else if str[0] == '$' {
		spls := strings.Split(str, "\r\n")
		strs = append(strs, spls[1])
	} else {
		spls := strings.Split(str, "\r\n")
		fmt.Println("spls: ", spls)
		for _, spl := range spls {
			//fmt.Println("for loop, idx: ", idx, "spl: ", spl)
			if spl == "" || spl[0] == '*' || spl[0] == '$' && (spl[0] != '+' && spl[0] != '-' && spl[0] != ':') {
				if  len(spl) > 1 && spl == "$-1" {
					strs = append(strs, "")
				}
			} else {
				strs = append(strs, spl)
			}
		}
	}

	return strs
}

func main()  {
	//str := "*2\r\n$3\r\nfoo\r\n$4\r\nbars\r\n"
	str := "$-1\r\n"
	//str := "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Foo\r\n-Bar\r\n"
	//str := "*3\r\n$3\r\nfoo\r\n$-1\r\n$3\r\nbar\r\n"
	//str := "-Error message\r\n"
	//str := "$3\r\nfoo\r\n"
	//str := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	//str := "$6\r\nfoobar\r\n"
	strs := parseString(str)
	for i, str := range strs {
		fmt.Println("index:, ", i, "str: ", str)
	}
}