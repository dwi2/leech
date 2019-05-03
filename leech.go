package main

import (
	"fmt"
	"leech/validpalindrome"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	actual := validpalindrome.IsPalindrome(input)
	// expect true
	fmt.Printf("%v\n", actual)
}
