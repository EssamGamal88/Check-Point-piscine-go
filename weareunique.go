package main

import "fmt"

func WeAreUnique(str1, str2 string) int {

	if len(str1) == 0 || len(str2) == 0 {
		return -1
	}

	unique := ""
	flag := true
	for _, ch := range str1 {
		for _, c := range str2 {
			if ch == c {
				flag = false
			}
		}
		if flag {
			unique += string(ch)
		}
		flag = true
	}
	for _, ch := range str2 {
		for _, c := range str1 {
			if ch == c {
				flag = false
			}
		}
		if flag {
			unique += string(ch)
		}
		flag = true
	}
	return len(unique)
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
