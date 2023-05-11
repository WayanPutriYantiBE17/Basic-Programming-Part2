package main

import "fmt"

func Palindrome(input string) bool {
	// your code here

	runeSlice := []rune(input)
	for i:=0;i<len(runeSlice)/2;i++{
		if runeSlice[i]!=runeSlice[len(runeSlice)-i-1]{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Palindrome("civic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
