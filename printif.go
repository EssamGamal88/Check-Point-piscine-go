package main

import "fmt"

func PrintIf(str string) string {
	if len(str) < 3 && str != "" {
		return "Invalid Input\n"
	}
	return "G\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
