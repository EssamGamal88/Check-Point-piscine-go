package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z') {
			return s
		}
		if i > 0 && c >= 'A' && c <= 'Z' && c >= s[i-1] && c <= s[i-1] {
			return s
		}
		Last := s[len(s)-1]
		if Last >= 'A' && Last <= 'Z' {
			return s
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i != 0 {
			if ch >= 'A' && ch <= 'Z' {
				result += "_"
			}
		}
		result += string(ch)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
