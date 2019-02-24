package main

import "fmt"

func main() {
	// 65 to 90 is the ASCII representation of alphabets A to Z
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 1; j <= 3; j++ {
			// %U is the unicode format (rune/code point) and %q is the character representation
			fmt.Printf("%U\t%q\n", i, i)
		}
	}
}
