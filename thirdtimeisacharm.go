package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	result := ""
	if str == "" || len(str) < 3 {
		return ""
	}
	for i, r := range str {
		if (i+1)%3 == 0 {
			result += string(r)
		}
	}
	return result
}

func main() {
	fmt.Println(ThirdTimeIsACharm("123456789"))
	fmt.Println(ThirdTimeIsACharm(""))
	fmt.Println(ThirdTimeIsACharm("a b c d e f"))
	fmt.Println(ThirdTimeIsACharm("12"))
}
