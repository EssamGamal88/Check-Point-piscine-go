package main

import (
	"fmt"
	"unicode"
)

func PrintMemory(arr [10]byte) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%02x", arr[i])
		if i != 9 {
			fmt.Print(" ")
		}
		if i == 3 || i == 7 {
			fmt.Println()
		}
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		r := rune(arr[i])
		if unicode.IsGraphic(r) {
			fmt.Printf("%c", r)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
